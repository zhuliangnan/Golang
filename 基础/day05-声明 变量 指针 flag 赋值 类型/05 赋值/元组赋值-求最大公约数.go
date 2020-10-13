//GCD -- 最大公约数
package main

import (
	"fmt"
)

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func main() {

	fmt.Println(gcd(20, 40))
}
