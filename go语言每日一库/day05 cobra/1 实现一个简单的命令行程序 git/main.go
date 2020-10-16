package main

/*
cobra是一个命令行程序库，可以用来编写命令行程序。
同时，它也提供了一个脚手架， 用于生成基于 cobra 的应用程序框架。非常多知名的开源项目使用了 cobra 库构建命令行，
如Kubernetes、Hugo、etcd等等等等。 本文介绍 cobra 库的基本使用和一些有趣的特性。
*/
import (
	"github.com/darjun/go-daily-lib/cobra/get-started/cmd"
)

func main() {
	cmd.Execute()
}
