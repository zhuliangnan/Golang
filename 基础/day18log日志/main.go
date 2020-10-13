package main

import "log"

/*
const (
	Ldate         = 1 << iota     //日期示例： 2009/01/23
	Ltime                         //时间示例: 01:23:23
	Lmicroseconds                 //毫秒示例: 01:23:23.123123.
	Llongfile                     //绝对路径和行号: /a/b/c/d.go:23
	Lshortfile                    //文件和行号: d.go:23.
	LUTC                          //日期时间转为0时区的
	LstdFlags     = Ldate | Ltime //Go提供的默认标准抬头信息
)
*/
func main() {
	log.Println("codesuger的博客", "http://www.codesuger.com/")
	log.Printf("codesuger的go语言教程:%s\n", "http://www.codesuger.com/categories/go语言教程")
}

//时间+文件名+源代码所在行号
func init() {
	//指定前缀
	log.SetPrefix("[codesuger]")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
