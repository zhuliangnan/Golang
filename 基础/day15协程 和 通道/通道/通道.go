package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 3) //带缓冲的通道

	//内置函数 len 返回未被读取的缓冲元素数量， cap 返回缓冲区大小
	fmt.Printf("len(c)=%d, cap(c)=%d\n", len(c), cap(c))

	fmt.Println("-----------开始------------")
	go func() {
		//defer 将一个方法延迟到包裹该方法的方法返回时执行
		defer fmt.Println("发送协程结束")

		for i := 0; i < 4; i++ {
			if i == 3 {
				fmt.Println("发送协程阻塞")
			}
			c <- i
			fmt.Printf("子协程正在发送:%d, len(c)=%d, cap(c)=%d\n", i, len(c), cap(c))
		}
		fmt.Println("发送协程阻塞放开")
	}()
	time.Sleep(time.Second) //延时1s 给子协程点时间让它执行完

	go func() {
		fmt.Println("-----------接受协程开始接受数据-----------")
		for i := 0; i < 4; i++ {
			if i == 3 {
				fmt.Println("接受协程阻塞")
			}
			n := <-c
			fmt.Printf("main协程接收到的数据:%d ,len(c)=%d, cap(c)=%d\n", n, len(c), cap(c))
		}

	}()

	time.Sleep(time.Second) //延时1s 给子协程点时间让它执行完
	fmt.Println("main协程结束")
}
