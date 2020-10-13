package main

import (
	"fmt"
	"sort"
	"sync"
)

func main() {

	ages := map[string]int{}
	ages["charlie"] = 23
	ages["jack"] = 16
	ages["tom"] = 19
	ages["yom"] = 27
	ages["bili"] = 20
	ages["alice"] = 14
	var names []string //声明一个字符串切片，存储map的key值
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names) //根据map的key排序
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}

	var scene sync.Map
	scene.Store("name", "coding3min")
	scene.Store("age", 11)

	v, ok := scene.Load("name")
	if ok {
		fmt.Println(v)
	}
	v, ok = scene.Load("age")
	if ok {
		fmt.Println(v)
	}

	scene.Delete("age")

	scene.Range(func(key, value interface{}) bool {
		fmt.Println("key:", key, ",value:", value)
		return true
	})

}
