package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

/*
godotenv库从.env文件中读取配置， 然后存储到程序的环境变量中
*/
func main() {

	//如果不想加载 可以直接导入  _ "github.com/joho/godotenv/autoload"
	//默认情况下是记载项目根目录下的环境配置.env文件
	//如果多个文件中存在同一个键，那么先出现的优先，后出现的不生效
	err := godotenv.Load("go语言每日一库/day10 godotenv/1 简单使用/.env")
	if err != nil {
		log.Fatal(err)
	}

	//调用os.Getenv("key")读取
	//os.Getenv是用来读取环境变量的
	fmt.Println(os.Getenv("GORoot"))
	fmt.Println("name: ", os.Getenv("name"))
	fmt.Println("age: ", os.Getenv("age"))
}
