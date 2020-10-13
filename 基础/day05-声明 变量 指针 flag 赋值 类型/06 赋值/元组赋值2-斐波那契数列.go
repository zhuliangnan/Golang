//斐波那契数列
package main

import (
	"fmt"
)

func fid(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}

func main() {

	fmt.Println(fid(10))
}
