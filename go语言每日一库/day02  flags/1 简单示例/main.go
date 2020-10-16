package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
)

/*
使用go-flags的一般步骤：

定义选项结构，在结构标签中设置选项信息。通过short和long设置短、长选项名字，description设置帮助信息。命令行传参时，短选项前加-，长选项前加--；
声明选项变量；
调用go-flags的解析方法解析。
*/

type Option struct {
	Verbose []bool `short:"v" long:"verbose" description:"Show verbose debug message"`
}

func main() {
	var opt Option
	flags.Parse(&opt)

	fmt.Println(opt.Verbose)
}
