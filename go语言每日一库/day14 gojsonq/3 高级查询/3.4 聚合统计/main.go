package main

import (
	"fmt"
	"github.com/thedevsaddam/gojsonq"
)

func main() {
	//对一些字段做简单的统计，计算和、平均数、最大、最小值等：
	gq := gojsonq.New().File("go语言每日一库/day14 gojsonq/2 数据源/data.json").From("items")

	fmt.Println("Total Count:", gq.Sum("count"))
	fmt.Println("Min Price:", gq.Min("price"))
	fmt.Println("Max Price:", gq.Max("price"))
	fmt.Println("Avg Price:", gq.Avg("price"))

	//聚合统计类的方法都不会修改当前节点的指向，所以JSONQ对象可以重复使用！
	//还可以对数据进行分组和排序：
	fmt.Println(gq.From("items").GroupBy("price").Get())
	gq.Reset()
	fmt.Println(gq.From("items").SortBy("price", "desc").Get())

}
