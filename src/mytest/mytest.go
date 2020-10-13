package mytest

import (
	"fmt"
)

func Tag(tag int) {
	switch tag {
	case 1:
		fmt.Println("Android")
	case 2:
		fmt.Println("Go")
	case 3:
		fmt.Println("Java")
	default:
		fmt.Println("C")

	}
}

func Add(a, b int) int {
	return a + b
}
