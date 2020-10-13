package main

import "fmt"

type people struct {
	name string
	age  int
}

func (p *people) sayHello() {
	p.age *= 2
	fmt.Printf("*%s的地址 %p \n", p.name, p)
}

/*func (p people) toString() {
	p.age *= 2
	fmt.Printf("%s的地址 %p \n",p.name, &p)
}*/
func main() {

	/*	p1 := people{"p1",3}
		p1.toString()
		fmt.Println("p1的age:",p1.age)

		p2 := people{"p2",4}
		p2.toString()
		fmt.Println("p2的age:",p2.age)*/

	p3 := people{"p3", 3}
	p3.sayHello()
	fmt.Println("p3的age:", p3.age)

	p4 := people{"p3", 4}
	p4.sayHello()
	fmt.Println("p4的age:", p4.age)

}
