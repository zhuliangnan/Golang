package main

import (
	"fmt"
	"time"
)

//chan + select 自动关闭goroutine
func main() {

	stop := make(chan bool)

	go func() {

		select {
		case <-stop:
			fmt.Println("监控退出了，停止")
			return
		default:
			fmt.Println("goroutine监控中...")
			time.Sleep(2 * time.Second)

		}
	}()

	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")
	stop <- true

	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)

}
