package main

import (
	"fmt"
)

func main() {

	fmt.Println("----------map的创建及基本操作---------")

	/*	ages := map[string]int{
		"alice" : 14,
		"charlie" : 23,
	}*/
	ages := map[string]int{}
	//ages := make(map[string]int)
	ages["alice"] = 14
	ages["charlie"] = 23
	for k, v := range ages {
		fmt.Printf("k=%s v=%d\n", k, v)
	}
	/*k ,ok := ages["alice"]
	fmt.Println(k,ok)*/

	ages["alice"] = 29
	fmt.Println(ages)

	delete(ages, "alice") // remove element ages["alice"]
	fmt.Println(ages)
	fmt.Println("----------map的注意点---------")

	/*	fmt.Println(ages["bob"]) // 0
		ages["bob"] = ages["bob"] + 1
		fmt.Println(ages["bob"]) // 1*/
	ages["bob"] += 1
	//更简单可以写成
	ages["bob"]++
	fmt.Println(ages["bob"])

	fmt.Println("----------未初始化的map---------")
	var names map[string]int
	fmt.Println(names == nil)
	fmt.Println(len(names) == 0)
	//names["alice"] = 12
	fmt.Println("----------初始化的map---------")
	sexs := make(map[string]int)
	fmt.Println(sexs == nil)
	fmt.Println(len(sexs) == 0)

}

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}
