package main

import "fmt"

// 1，养父类
type AdoptedFather struct {
}

// 2, 养父的遗产
func (af AdoptedFather) Show() {
	println("I'm your Father! You succeeded in inheriting my legacy.")
}

// 3, 父类
type AdoptedChild struct {
	AdoptedFather
}

type people struct {
	name string
}

/*func (p people) toString() {
	fmt.Println(p.name)
	fmt.Printf("p的地址 %p \n", &p)
}*/
func (p *people) sayHello() {
	fmt.Printf("Hello! %v \n", p.name)
	fmt.Printf("*p的地址 %p \n", p)
}

func main() {
	/*	child := AdoptedChild{} // 干儿子
		child.Show()            // 继承干爹的遗产*/

	/*	p1 := people{"p1"}
		p1.toString()
		p2 := people{"p2"}
		p2.toString()*/

	p3 := people{"p3"}
	p3.sayHello()
	p4 := &people{"p4"}
	p4.sayHello() //系统帮你转成(*people).sayHello

	/*	p1 := &people{"coding3min"}
		p1.sayHello()
		p2 := people{"codesuger"}
		p2.sayHello()

	*/

}
