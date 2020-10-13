package main

import "fmt"

func add(x, y, i, k int) func() (int, int, int, int) {
	k++
	p := 0
	fmt.Println("p  k  初始化执行")
	return func() (int, int, int, int) {
		p++
		i++
		fmt.Println("i,k,p,x+y")
		return i, k, p, x + y
	}
}

func main() {
	fmt.Println("--------result--------")
	result := add(1, 2, 0, 0)
	fmt.Println(result())
	fmt.Println(result())
	fmt.Println(result())
	fmt.Println("--------result2--------")
	result2 := add(3, 4, 0, 0)
	fmt.Println(result2())

}
