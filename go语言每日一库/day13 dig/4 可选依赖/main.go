package main

/*
如果找不到对应的依赖，那么相应的对象无法创建成功，
调用Invoke时也会返回错误。有些依赖不是必须的，dig也提供了一种方式将依赖设置为可选的：
type Config struct {
  dig.In

  Redis *RedisConfig `optional:"true"`
  MySQL *MySQLConfig
}
*/
import (
	"fmt"
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

type Config struct {
	dig.In

	Redis *RedisConfig `optional:"true"`
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

/*
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
*/
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
	if config.Redis == nil {
		fmt.Println("no redis config")
	}
}

func main() {
	container := dig.New()

	container.Provide(InitOption)
	container.Provide(InitConfig)
	container.Provide(InitMySQLConfig)

	container.Invoke(PrintInfo)
}
