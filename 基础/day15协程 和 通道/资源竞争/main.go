package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	count int32
	wg    sync.WaitGroup
)

func main() {
	wg.Add(2)
	go incCount()
	go incCount()
	wg.Wait()
	fmt.Println(count)
}

func incCount() {
	defer wg.Done()
	for i := 0; i < 2; i++ {
		value := count
		//当前goroutine暂停的意思，退回执行队列，让其他等待的goroutine运行
		runtime.Gosched()
		value++
		count = value
	}
}
