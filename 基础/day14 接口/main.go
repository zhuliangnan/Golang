package main

import "fmt"

type Animal interface {
	run() string
	fly() string
	eat() string
}

//定义结构体
type Dog struct {
	name string
	age  int
}

func (dog Dog) fly() string {
	fmt.Println("Dog's name is", dog.name, "He cannot fly.Because he is dog,dog cannot fly ")
	return ""
}

//实现接口
func (dog Dog) run() string {
	fmt.Println("Dog's name is", dog.name, "He can run. He is", dog.age, "岁了")
	return ""
}
func (dog Dog) eat() string {
	fmt.Println("Dog's name is", dog.name, "He can eat. He is", dog.age, "岁了")
	return ""
}

func main() {
	var animal Animal
	dog := Dog{"dog", 12}
	animal = dog
	animal.run()
	animal.eat()
	animal.fly()
}
