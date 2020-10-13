package main

import (
	"errors"
	"fmt"
)

func Add(a, b int) (sub int, err error) {
	if a < 0 || b < 0 {
		err = errors.New("a，b不同为负数") // 这里的errors
		return 0, err
	}
	return a + b, nil
}

func Cacule(n int) (sum int) {
	if n > 0 {
		return Cacule(n-1) * n
	}
	return 1
}

func main() {
	fmt.Println(Add(-1, 2))

	fmt.Println(Cacule(10))
}
