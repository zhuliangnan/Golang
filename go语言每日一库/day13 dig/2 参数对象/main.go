package main

/*
有时候，创建对象有很多依赖，或者编写函数时有多个参数依赖。如果将这些依赖都作为参数传入，那么代码将变得非常难以阅读：
dig支持将所有参数打包进一个对象中，唯一需要的就是将dig.In内嵌到该类型中：
type Params {
  dig.In

  Arg1 *Arg1
  Arg2 *Arg2
  Arg3 *Arg3
  Arg4 *Arg4
}

container.Provide(func (params Params) *Object {
  // ...
})
内嵌了dig.In之后，dig会将该类型中的其它字段看成Object的依赖，创建Object类型的对象时，会先将依赖的Arg1/Arg2/Arg3/Arg4创建好。
*/

import (
	"fmt"
	"log"

	"github.com/jessevdk/go-flags"
	"go.uber.org/dig"
	"gopkg.in/ini.v1"
)

type Option struct {
	ConfigFile string `short:"c" long:"config" description:"Name of config file."`
}

type RedisConfig struct {
	IP   string
	Port int
	DB   int
}

type MySQLConfig struct {
	IP       string
	Port     int
	User     string
	Password string
	Database string
}

//参数对象
type Config struct {
	dig.In

	Redis *RedisConfig
	MySQL *MySQLConfig
}

func InitOption() (*Option, error) {
	var opt Option
	_, err := flags.Parse(&opt)

	return &opt, err
}

func InitConfig(opt *Option) (*ini.File, error) {
	cfg, err := ini.Load(opt.ConfigFile)
	return cfg, err
}

func InitRedisConfig(cfg *ini.File) (*RedisConfig, error) {
	port, err := cfg.Section("redis").Key("port").Int()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	db, err := cfg.Section("redis").Key("db").Int()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &RedisConfig{
		IP:   cfg.Section("redis").Key("ip").String(),
		Port: port,
		DB:   db,
	}, nil
}

func InitMySQLConfig(cfg *ini.File) (*MySQLConfig, error) {
	port, err := cfg.Section("mysql").Key("port").Int()
	if err != nil {
		return nil, err
	}

	return &MySQLConfig{
		IP:       cfg.Section("mysql").Key("ip").String(),
		Port:     port,
		User:     cfg.Section("mysql").Key("user").String(),
		Password: cfg.Section("mysql").Key("password").String(),
		Database: cfg.Section("mysql").Key("database").String(),
	}, nil
}

func PrintInfo(config Config) {
	fmt.Println("=========== redis section ===========")
	fmt.Println("redis ip:", config.Redis.IP)
	fmt.Println("redis port:", config.Redis.Port)
	fmt.Println("redis db:", config.Redis.DB)

	fmt.Println("=========== mysql section ===========")
	fmt.Println("mysql ip:", config.MySQL.IP)
	fmt.Println("mysql port:", config.MySQL.Port)
	fmt.Println("mysql user:", config.MySQL.User)
	fmt.Println("mysql password:", config.MySQL.Password)
	fmt.Println("mysql db:", config.MySQL.Database)
}

func main() {
	container := dig.New()

	container.Provide(InitOption)
	container.Provide(InitConfig)
	container.Provide(InitRedisConfig)
	container.Provide(InitMySQLConfig)

	err := container.Invoke(PrintInfo)
	if err != nil {
		log.Fatal(err)
	}
}
