package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	//使用一个逻辑处理器
	runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			fmt.Println("A:", i)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			fmt.Println("B:", i)
		}
	}()

	wg.Wait()
	//不要误认为是顺序执行，这里之所以顺序输出的原因，是因为我们的goroutine执行时间太短暂了，还没来得及切换到第2个goroutine，第1个goroutine就完成了

}
