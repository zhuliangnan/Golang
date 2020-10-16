package main

import (
	"encoding/json"
	"fmt"
	"github.com/thedevsaddam/gojsonq"
)

/*
gojsonq的独特之处在于，它可以像 SQL 一样进行条件查询，可以选择返回哪些字段，可以做一些聚合统计。
*/
func main() {
	//From方法，这个方法的作用是将当前节点移动到指定位置  最后必须要调用Get()，它组合所有条件后执行这个查询，返回结果。
	r := gojsonq.New().File("go语言每日一库/day14 gojsonq/2 数据源/data.json").From("items").Select("id", "name").Get()
	//用json.MarshalIndent()对输出做了一些美化。
	data, _ := json.MarshalIndent(r, "", "  ")
	//data, _ := json.Marshal(r)
	fmt.Println(string(data))
}
