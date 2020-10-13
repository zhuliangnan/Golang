package main

import "fmt"

func sumNumber(paras ...int) int {
	var sum int
	for _, v := range paras {
		sum += v
	}
	return sum

}

func main() {
	fmt.Println(sumNumber(1))             // 1
	fmt.Println(sumNumber(1, 2, 3, 4, 5)) // 15
	//在上面的代码中 , 调用者隐式的创建一个数组，并将原始参数复制到数组中，再把数组的一 个切片作为参数传给被调函数

	//如果原始参数已经是切片类型，我们该如何传递给sum？只需 在最后一个参数后加上省略符
	values := []int{1, 2, 3, 4, 5}
	fmt.Println(sumNumber(values...)) // 15
}
