package main

import "fmt"

type Dog struct {
	name string
	age  int
}

func (dog Dog) String() string {
	return fmt.Sprintf("%s---%v", dog.name, dog.age)

}
func main() {
	fmt.Println(Dog{"dog", 12})
}
