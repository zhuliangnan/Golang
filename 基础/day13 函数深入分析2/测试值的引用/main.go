package main

import "fmt"

func main() {
	i := 10 //整形变量 i
	//p := &i //指向整型变量 i 的指针ip,包含了 i 的内存地址
	fmt.Printf("main中i的值为：%v，i 的内存地址为：%v\n", i, &i)
	searchByP(&i)
	fmt.Printf("main中i的值为：%v，i 的内存地址为：%v\n", i, &i)
}

func searchByP(i *int) {
	fmt.Printf("指向i的指针的内存地址为：%v\n", &i)
	*i = 11
}
