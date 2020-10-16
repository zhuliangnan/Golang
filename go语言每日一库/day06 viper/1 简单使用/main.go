package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

/*
 viper 是一个配置解决方案，拥有丰富的特性：

支持 JSON/TOML/YAML/HCL/envfile/Java properties 等多种格式的配置文件；
可以设置监听配置文件的修改，修改时自动加载新的配置；
从环境变量、命令行选项和io.Reader中读取配置；
从远程配置系统中读取和监听修改，如 etcd/Consul；
代码逻辑中显示设置键值。
*/
func init() {
	//pflag.Int("redis.port", 8381, "Redis port to connect")

	// 绑定命令行
	//viper.BindPFlags(pflag.CommandLine)

	// 绑定环境变量
	//viper.AutomaticEnv()
}

func main() {
	//调用pflag.Parse解析命令行
	//pflag.Parse()

	//设置文件名  设置文件名时不要带后缀；
	//
	///go语言每日一库/day06 viper/1 简单使用/config
	viper.SetConfigName("go语言每日一库/day06 viper/1 简单使用/config")
	//配置类型
	viper.SetConfigType("toml")
	//搜索路径  搜索路径可以设置多个，viper 会根据设置顺序依次查找；
	viper.AddConfigPath(".")

	//默认值可以调用viper.SetDefault设置。
	viper.SetDefault("redis.port", 6381)
	//开始读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("read config failed: %v", err)
	}

	fmt.Println("-------Set值的优先级 Set》命令行》环境变量》配置文件-----------")
	viper.Set("redis.port", 5381)

	//如果一个键没有通过viper.Set显示设置值，那么获取时将尝试从命令行选项中读取。 如果有，优先使用
	//首先在init方法中定义选项，并且调用viper.BindPFlags绑定选项到配置中：
	//然后，在main方法开头处调用pflag.Parse解析选项。
	//$ ./main.exe --redis.port 9381

	//如果前面都没有获取到键值，将尝试从环境变量中读取。我们既可以一个个绑定，也可以自动全部绑定。
	//
	//在init方法中调用AutomaticEnv方法绑定全部环境变量：

	fmt.Println("-------Get方法的使用-----------")
	fmt.Println(viper.Get("app_name"))
	fmt.Println(viper.Get("log_level"))

	//viper 获取值时使用section.key的形式，即传入嵌套的键名；
	fmt.Println("mysql ip: ", viper.Get("mysql.ip"))
	fmt.Println("mysql port: ", viper.Get("mysql.port"))
	fmt.Println("mysql user: ", viper.Get("mysql.user"))
	fmt.Println("mysql password: ", viper.Get("mysql.password"))
	fmt.Println("mysql database: ", viper.Get("mysql.database"))

	fmt.Println("redis ip: ", viper.Get("redis.ip"))
	fmt.Println("redis port: ", viper.Get("redis.port"))

	fmt.Println("-------GetType方法的使用-----------")
	//如果指定的键不存在或类型不正确，GetType方法返回对应类型的零值。
	fmt.Println("protocols: ", viper.GetStringSlice("server.protocols"))
	fmt.Println("ports: ", viper.GetIntSlice("server.ports"))
	fmt.Println("timeout: ", viper.GetDuration("server.timeout"))

	fmt.Println("mysql ip: ", viper.GetString("mysql.ip"))
	fmt.Println("mysql port: ", viper.GetInt("mysql.port"))

	//如果要判断某个键是否存在，使用IsSet方法
	if viper.IsSet("redis.port") {
		fmt.Println("redis.port is set")
	} else {
		fmt.Println("redis.port is not set")
	}

	//GetStringMap和GetStringMapString直接以 map 返回某个键下面所有的键值对，
	//前者返回map[string]interface{}，
	//后者返回map[string]string。
	//AllSettings以map[string]interface{}返回所有设置。
	fmt.Println("mysql settings: ", viper.GetStringMap("mysql"))
	fmt.Println("redis settings: ", viper.GetStringMap("redis"))
	fmt.Println("all settings: ", viper.AllSettings())

}
