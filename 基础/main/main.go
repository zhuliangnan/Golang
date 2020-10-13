package main

import (
	"fmt"
	"strconv"
)

func main() {

	var aInt int = 17
	fmt.Printf("转 float64 %f  \n", float64(aInt))
	fmt.Printf("转 string %v  \n", strconv.Itoa(aInt))
	fmt.Printf("转 float64 %f  \n", float64(aInt))
	fmt.Printf("南京南 %f ", float64(30))
	//各种类型转字符串
	var resString string
	resString = fmt.Sprintf("%d %v %v", 40, "code", true)

	re := fmt.Sprintf("%g", 123.002)
	fmt.Println(re)

	fmt.Println(resString)

	// string  to bytes
	resBety := []byte("ASQWD")
	// bytes to string
	String := string(resBety)

	fmt.Println(resBety, String)
}
