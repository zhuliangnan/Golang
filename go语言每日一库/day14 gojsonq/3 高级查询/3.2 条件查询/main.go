package main

import (
	"fmt"
	"github.com/thedevsaddam/gojsonq"
)

func main() {
	gq := gojsonq.New().File("go语言每日一库/day14 gojsonq/2 数据源/data.json")

	r := gq.From("items").Select("id", "name").
		Where("id", "=", 1).OrWhere("id", "=", 2).Get()
	fmt.Println(r)

	gq.Reset()

	r = gq.From("items").Select("id", "name", "count").
		Where("count", ">", 1).Where("price", "<", 100).Get()
	fmt.Println(r)
}
