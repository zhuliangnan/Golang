package main

import (
	"github.com/spf13/viper"
	"log"
)

/*
有时候，我们想要将程序中生成的配置，或者所做的修改保存下来。viper 提供了接口！

WriteConfig：将当前的 viper 配置写到预定义路径，如果没有预定义路径，返回错误。将会覆盖当前配置；
SafeWriteConfig：与上面功能一样，但是如果配置文件存在，则不覆盖；
WriteConfigAs：保存配置到指定路径，如果文件存在，则覆盖；
SafeWriteConfig：与上面功能一样，但是入股配置文件存在，则不覆盖。
*/

//下面我们通过程序生成一个config.toml配置：

func main() {
	viper.SetConfigName("go语言每日一库/day06 viper/4 保存配置/config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	viper.Set("app_name", "awesome web")
	viper.Set("log_level", "DEBUG")
	viper.Set("mysql.ip", "127.0.0.1")
	viper.Set("mysql.port", 3306)
	viper.Set("mysql.user", "root")
	viper.Set("mysql.password", "123456")
	viper.Set("mysql.database", "awesome")

	viper.Set("redis.ip", "127.0.0.1")
	viper.Set("redis.port", 6382)

	//如果文件路径存在文件 则覆盖
	err := viper.WriteConfig()
	if err != nil {
		log.Fatal("write config failed: ", err)
	}
	//如果文件路径存在文件 则不覆盖
	/*err := viper.SafeWriteConfig()
	if err != nil {
		log.Fatal("write config failed: ", err)
	}*/
}
