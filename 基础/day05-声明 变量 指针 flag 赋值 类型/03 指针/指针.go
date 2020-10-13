package main

import (
	"fmt"
)

func main() {

	x := 1

	p := &x

	fmt.Println(p)
	fmt.Println(*p)
	*p = 2
	fmt.Println(x)
}
