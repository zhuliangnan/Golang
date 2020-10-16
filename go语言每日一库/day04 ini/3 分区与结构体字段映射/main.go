package main

/*
定义结构变量，加载完配置文件后，调用MapTo将配置项赋值到结构变量的对应字段中
*/
import (
	"fmt"
	"gopkg.in/ini.v1"
)

type Config struct {
	AppName  string `ini:"app_name"`
	LogLevel string `ini:"log_level"`

	MySQL MySQLConfig `ini:"mysql"`
	Redis RedisConfig `ini:"redis"`
}

type MySQLConfig struct {
	IP       string `ini:"ip"`
	Port     int    `ini:"port"`
	User     string `ini:"user"`
	Password string `ini:"password"`
	Database string `ini:"database"`
}

type RedisConfig struct {
	IP   string `ini:"ip"`
	Port int    `ini:"port"`
}

func main() {
	cfg, err := ini.Load("go语言每日一库/day04 ini/3 分区与结构体字段映射/my.ini")
	if err != nil {
		fmt.Println("load my.ini failed: ", err)
	}

	c := Config{}
	//用MapTo将配置项赋值到结构变量的对应字段中
	cfg.MapTo(&c)

	fmt.Println(c)
}
