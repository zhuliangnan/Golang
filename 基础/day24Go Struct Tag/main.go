package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type User struct {
	Name string `name`
	Age  int    `age`
}

//Json字符串转User对象
//里主要利用的就是User这个结构体对应的字段Tag，
//json解析的原理就是通过反射获得每个字段的tag，然后把解析的json对应的值赋给他们。
func main() {
	var u User
	h := `{"name" : "张三" , "age" : 15}`
	err := json.Unmarshal([]byte(h), &u)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(u)
	}

	//把结构体转换为Json字符串
	newJson, err := json.Marshal(&u)
	fmt.Println((string(newJson)))

	fmt.Println("--------反射获取字段Tag-------")
	//字段的Tag是标记到字段上的，所以我们可以通过先获取字段，然后再获取字段上的Tag。
	t := reflect.TypeOf(u)

	for i := 0; i < t.NumField(); i++ {
		sf := t.Field(i)
		fmt.Println(sf.Tag)
	}

	fmt.Println("--------字段Tag的键值对-------")
	//很多时候我们的一个Struct不止具有一个功能，比如我们需要JSON的互转、
	//还需要BSON以及ORM解析的互转，所以一个字段可能对应多个不同的Tag，以便满足不同的功能场景。
	var p Person
	pt := reflect.TypeOf(p)

	for i := 0; i < pt.NumField(); i++ {
		sf := pt.Field(i)
		fmt.Println(sf.Tag.Get("json"))

	}

	var p2 Person2
	pt2 := reflect.TypeOf(p2)

	for i := 0; i < pt2.NumField(); i++ {
		sf2 := pt2.Field(i)
		fmt.Println(sf2.Tag.Get("bson"))

	}

}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
type Person2 struct {
	Name string `json:"name" bson:"b_name"`
	Age  int    `json:"age" bson:"b_age"`
}
