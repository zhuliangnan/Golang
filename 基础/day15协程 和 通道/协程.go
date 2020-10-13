package main

import (
	"fmt"
	"time"
)

func add(x, y int) int {

	fmt.Println("计算结果是:", x+y)
	return x + y
}

func main() {

	fmt.Println("开始执行main")
	for i := 0; i < 10; i++ {
		go add(i, i)
	}
	time.Sleep(time.Second)
	fmt.Println("main执行结束的操作")

}
