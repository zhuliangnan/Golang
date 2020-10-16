package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	//先读取环境变量GODAILYLIB_ENV -- 先读取到的优先
	env := os.Getenv("GODAILYLIB_ENV")
	if env == "" {
		env = "development"
	}

	//读取对应的.env. + env
	err := godotenv.Load(".env." + env)
	if err != nil {
		log.Fatal(err)
	}

	//最后读取默认的.env文件
	err = godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("name: ", os.Getenv("name"))
	fmt.Println("version: ", os.Getenv("version"))
	fmt.Println("database: ", os.Getenv("database"))

	/*
		# 默认是开发环境
		$ go run main.go
		name:  awesome web
		version:  0.0.1
		database:  sqlite3

		# 设置为生成环境
		$ GODAILYLIB_ENV=production go run main.go
		name:  awesome web
		version:  0.0.1
		database:  mysql
	*/
}
