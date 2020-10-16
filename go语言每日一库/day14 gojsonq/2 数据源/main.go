package main

import (
	"fmt"
	"github.com/thedevsaddam/gojsonq"
)

/*
除了从字符串中加载，jsonq还允许从文件和io.Reader中读取内容。分别使用JSONQ对象的File和Reader方法：
*/
func main() {
	gq := gojsonq.New().File("go语言每日一库/day14 gojsonq/2 数据源/data.json")

	fmt.Println(gq.Find("items.[1].price"))

	/*	file, err := os.OpenFile("go语言每日一库/day14 gojsonq/2 数据源/data.json", os.O_RDONLY, 0666)
		if err != nil {
			log.Fatal(err)
		}

		gq := gojsonq.New().Reader(file)

		fmt.Println(gq.Find("items.[1].price"))*/
}
