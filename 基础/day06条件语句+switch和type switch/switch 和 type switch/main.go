package main

import "fmt"

func main() {
	name := ""

	switch name {
	case "zhuzhu":
		fmt.Println("zhuzhu")
	case "nanna":
		fmt.Println("nanan")
	case "liangliang":
		fmt.Println("liang")
	default:
		fmt.Println("默认")

	}

	number := 70

	//如果没有一个是匹配的，就执行default后的语句。
	//注意switch后可以跟空，如上
	//switch {
	//这样case就必须是表达式。
	switch {
	case number >= 90:
		fmt.Println("优秀")
	case number >= 80:
		fmt.Println("良好")
	case number >= 60:
		fmt.Println("凑合")
	default:
		fmt.Println("太搓了")
	}

	//有一个流传于坊间的神秘玩法，可以用switch语句来判断传入变量的类型，然后做一些羞羞的事情。
	//x是一个未知类型的变量，switch t := x.(type) 用这个方式来赋值，t就是有确定类型的变量。
	//type switch只能用于接口的变量。

	var x interface{} //定义接口类型
	x = "abc"         //实现接口类型  实现为 string
	switch t := x.(type) {
	case int:
		fmt.Println(t)
	case float64:
		fmt.Println("float64", t)
	default:
		fmt.Println("未知类型")

	}

}
