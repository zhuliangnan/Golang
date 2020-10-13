package main

import (
	"fmt"
	"time"
)

type T int

func IsClosed(ch <-chan T) bool {

	select {
	case <-ch:
		return true
	default:
	}
	return false
}

func SafeSend(ch chan T, value T) (sended bool) {
	sended = true
	defer func() {
		/*
			Recover()用法是：将Recover()写在defer中，
			并且在可能发生panic的地方之前，
			先调用此defer的东西（让系统方法域结束时，有代码要执行。）当程序遇到panic的时候（当然，也可以正常的调用出现的异常情况），
			系统将跳过后面的代码，进入defer，如果defer函数中有recover()，则返回捕获到的panic的值。
		*/
		if recover() != nil {
			sended = false
		}
	}()

	ch <- value   // panic if ch is closed
	return sended // <=> closed = false; return

}
func main() {

	/*	c := make(chan T)
		fmt.Println(IsClosed(c))
		close(c)
		fmt.Println(IsClosed(c))
	*/

	c := make(chan T)
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("收到", <-c)
		}

	}()
	fmt.Println(SafeSend(c, 3))
	fmt.Println(SafeSend(c, 4))
	close(c)
	fmt.Println(SafeSend(c, 5))

	time.Sleep(time.Second)

}
