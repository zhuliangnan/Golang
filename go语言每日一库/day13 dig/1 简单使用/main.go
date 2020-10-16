package main

import (
	"fmt"

	"github.com/jessevdk/go-flags"
	"go.uber.org/dig"
	"gopkg.in/ini.v1"
)

/*
今天我们来介绍 Go 语言的一个依赖注入（DI）库——dig。
dig 是 uber 开源的库。Java 依赖注入的库有很多，相信即使不是做 Java 开发的童鞋也听过大名鼎鼎的 Spring。
相比庞大的 Spring，dig 很小巧，实现和使用都比较简洁。

$ go get go.uber.org/dig
$ go get gopkg.in/ini.v1
$ go get github.com/jessevdk/go-flags


dig库帮助开发者管理这些对象的创建和维护，每种类型的对象会创建且只创建一次。
dig库使用的一般流程：
* 创建一个容器：dig.New；
* 为想要让dig容器管理的类型创建构造函数，构造函数可以返回多个值，这些值都会被容器管理；
* 使用这些类型的时候直接编写一个函数，将这些类型作为参数，然后使用container.Invoke执行我们编写的函数。
*/
type Option struct {
	ConfigFile string `short:"c" long:"config" description:"Name of config file."`
}

func InitOption() (*Option, error) {
	var opt Option
	_, err := flags.Parse(&opt)

	return &opt, err
}

func InitConf(opt *Option) (*ini.File, error) {
	cfg, err := ini.Load(opt.ConfigFile)
	return cfg, err
}

func PrintInfo(cfg *ini.File) {
	fmt.Println("App Name:", cfg.Section("").Key("app_name").String())
	fmt.Println("Log Level:", cfg.Section("").Key("log_level").String())
}

func main() {
	//创建一个容器：dig.New；
	container := dig.New()

	//为想要让dig容器管理的类型创建构造函数，构造函数可以返回多个值，这些值都会被容器管理；
	container.Provide(InitOption)
	container.Provide(InitConf)

	//使用这些类型的时候直接编写一个函数，将这些类型作为参数，然后使用container.Invoke执行我们编写的函数。
	container.Invoke(PrintInfo)
}
