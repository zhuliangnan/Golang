package main

import (
	"fmt"
	"github.com/thedevsaddam/gojsonq"
)

func main() {
	gq := gojsonq.New().File("go语言每日一库/day14 gojsonq/2 数据源/data.json")

	//第一次查询时返回前 3 条内容，第二次查询时返回接下来的 3 条记录。我们可以使用JSONQ对象的Offset和Limit方法来指定偏移和返回的条目数：
	r1 := gq.From("items").Select("id", "name").Offset(0).Limit(3).Get()
	fmt.Println("First Page:", r1)

	gq.Reset()

	r2 := gq.From("items").Select("id", "name").Offset(3).Limit(3).Get()
	fmt.Println("Second Page:", r2)
}
