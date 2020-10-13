package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	count int32
	wg    sync.WaitGroup
	mutex sync.Mutex
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
		mutex.Lock()
		value := count
		runtime.Gosched()
		value++
		count = value
		mutex.Unlock()
	}

}
