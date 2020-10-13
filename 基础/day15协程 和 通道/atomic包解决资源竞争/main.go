package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
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
		//读取int32类型变量的值
		value := atomic.LoadInt32(&count)
		runtime.Gosched()
		value++
		//修改int32类型变量的值
		atomic.StoreInt32(&count, value)
	}
}
