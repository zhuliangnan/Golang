package main

/*
前面我们说过，dig默认只会为每种类型创建一个对象。如果要创建某个类型的多个对象怎么办呢？可以为对象命名！

调用容器的Provide方法时，可以为构造函数的返回对象命名，这样同一个类型就可以有多个对象了。
*/

import (
	"fmt"

	"go.uber.org/dig"
)

type User struct {
	Name string
	Age  int
}

func NewUser(name string, age int) func() *User {
	return func() *User {
		return &User{name, age}
	}
}

type UserParams struct {
	dig.In

	User1 *User `name:"dj"`
	User2 *User `name:"dj2"`
}

func PrintInfo(params UserParams) error {
	fmt.Println("User 1 ===========")
	fmt.Println("Name:", params.User1.Name)
	fmt.Println("Age:", params.User1.Age)

	fmt.Println("User 2 ===========")
	fmt.Println("Name:", params.User2.Name)
	fmt.Println("Age:", params.User2.Age)
	return nil
}

func main() {
	container := dig.New()

	container.Provide(NewUser("dj", 18), dig.Name("dj"))
	container.Provide(NewUser("dj2", 18), dig.Name("dj2"))

	container.Invoke(PrintInfo)
}
