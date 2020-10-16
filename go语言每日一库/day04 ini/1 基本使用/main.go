package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"log"
)

/*
ini 是 Windows 上常用的配置文件格式。MySQL 的 Windows 版就是使用 ini 格式存储配置的。
go-ini是 Go 语言中用于操作 ini 文件的第三方库。


首先调用ini.Load加载文件，得到配置对象cfg；
然后以分区名调用配置对象的Section方法得到对应的分区对象section，默认分区的名字为""，也可以使用ini.DefaultSection；
以键名调用分区对象的Key方法得到对应的配置项key对象；
由于文件中读取出来的都是字符串，key对象需根据类型调用对应的方法返回具体类型的值使用，如上面的String、MustInt方法。
*/
func main() {
	cfg, err := ini.Load("go语言每日一库/day04 ini/1 基本使用/my.ini")
	if err != nil {
		log.Fatal("Fail to read file: ", err)
	}
	//以通过Sections方法获取所有分区，SectionStrings()方法获取所有分区名。
	names := cfg.SectionStrings()
	fmt.Println("names: ", names)

	fmt.Println("App Name:", cfg.Section("").Key("app_name").String())
	fmt.Println("Log Level:", cfg.Section("").Key("log_level").String())

	fmt.Println("MySQL IP:", cfg.Section("mysql").Key("ip").String())
	mysqlPort, err := cfg.Section("mysql").Key("port").Int()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MySQL Port:", mysqlPort)
	fmt.Println("MySQL User:", cfg.Section("mysql").Key("user").String())
	fmt.Println("MySQL Password:", cfg.Section("mysql").Key("password").String())
	fmt.Println("MySQL Database:", cfg.Section("mysql").Key("database").String())

	fmt.Println("Redis IP:", cfg.Section("redis").Key("ip").String())
	/*	redisPort, err := cfg.Section("redis").Key("port").Int()
		if err != nil {
			log.Fatal(err)
		}*/
	//这个方法只返回一个值。 同时它接受可变参数，如果类型无法转换，取参数中第一个值返回，并且该参数设置为这个配置的值，下次调用返回这个值
	fmt.Println("redis Port:", cfg.Section("redis").Key("port").MustInt(6381))
	//fmt.Println("Redis Port:", redisPort)

	//可以使用占位符%(name)s表示用之前已定义的键name的值来替换，这里的s表示值为字符串类型
	fmt.Println("Clone url:", cfg.Section("redis.sub").Key("CLONE_URL").String())
}
