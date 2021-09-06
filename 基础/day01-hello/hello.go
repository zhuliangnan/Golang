package main

import (
	"fmt"
	"time"
)

func main() {

	num2 := 0
	num3 := 0

	go func() {
		for i := 0; i < 100; i++ {
			num2++
		}
	}()
	time.Sleep(time.Second)
	fmt.Println(num2)
	go func() {
		num3++
	}()
	time.Sleep(time.Second)
	fmt.Println(num3)

	num := 0
	go func() {
		for {
			*(&num) = *(&num) + 1

			//if num/100 == 0 {
			//fmt.Print("*")
			//fmt.Println("num", num)
			//}

			/*if num == 10 {
				break
			}*/
			//runtime.Gosched()
		}
	}()
	time.Sleep(time.Second * 1)
	fmt.Println("num++++++++++++++++++++++++++++++++", num)
}
