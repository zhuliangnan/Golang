package main

import "fmt"

func main() {
	dataChan := make(chan int, 3)
	startChan := make(chan string, 1)
	overChan := make(chan string, 2)

	go func() {
		<-startChan
		fmt.Println("start receive data ")
		for {
			if elem, ok := <-dataChan; ok {
				fmt.Printf("%v\n", elem)
			} else {
				break
			}
		}
		fmt.Println("reciver over")
		overChan <- "rec over"
	}()

	go func() {
		for i := 0; i < 3; i++ {
			dataChan <- i
			fmt.Println("发送数据：", i)
		}

		fmt.Println("send data over")
		//在接收之前关闭通道
		close(dataChan)
		fmt.Println("dataChan  closed")

		//开始接收
		startChan <- "begin"

		overChan <- "send data over"
	}()

	<-overChan
	<-overChan
	fmt.Println("main over")

}
