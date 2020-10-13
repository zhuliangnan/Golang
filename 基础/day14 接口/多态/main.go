package main

import "fmt"

type Animal interface {
	fly() string
}

//定义结构体
type Dog struct {
	name string
	age  int
}

//定义结构体
type Bird struct {
	name string
	age  int
}

//一个以接口为参数的函数
func animalFly(animal Animal) {
	animal.fly()
}

func (dog Dog) fly() string {
	fmt.Println("Dog's name is", dog.name, "He cannot fly.Because he is dog,dog cannot fly ")
	return ""
}

func (bird Bird) fly() string {
	fmt.Println("Bird's name is", bird.name, "she can fly.Because he is bird,bird can fly ")
	return ""
}

func main() {
	dog := Dog{"dog", 12}
	bird := Bird{"bird", 10}
	animalFly(dog)
	animalFly(bird)

}
