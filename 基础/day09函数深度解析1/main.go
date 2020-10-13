package main

import (
	"fmt"
)

// Tip :函数没有返回值的时候，要么是只需要打印结果，要么是只做单元测试，除了这两种情况，没有返回值的函数就是做了很多事情的你没有和老板汇报一样，没有任何意义！

//一个返回值
func funcReturnOne() int {
	return 1
}

//返回值有名称
func funcReturnName() (res int) {
	//var res  int  省掉了
	res = 1 + 2
	return res
}

//多返回值
func funcReturnMore() (int, string, float64) {
	i := 1
	str := "zhu"
	fl := 23.02
	return i, str, fl
}

//值传递
func noChange(a, b int) {
	tmp := a
	a = b
	b = tmp
}

// 引用传递，参数加*号代表指针
func change(a, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}

//函数作为参数
func functionValue(a, b int, do func(int, int) int) {
	fmt.Println(do(a, b))
}
func add(a, b int) int {
	return a + b
}
func sub(a, b int) int {
	return a - b
}

func main() {
	i := funcReturnOne()
	fmt.Println(i)

	i2 := funcReturnName()
	fmt.Println(i2)

	a1, str, fl := funcReturnMore()
	fmt.Printf("%d %s %f \n", a1, str, fl)

	//值传递
	a, b := 1, 2
	fmt.Printf("原值 a:%v,b:%v \n", a, b)
	noChange(a, b)
	//值传递，并没有修改原值
	fmt.Printf("值传递后 a:%v,b:%v \n", a, b)

	// 引用传递，参数加*号代表指针
	aa, bb := 20, 30
	change(&aa, &bb)
	fmt.Printf("引用传递之后的值  aa : %v  bb : %v  \n", aa, bb)

	//上面的例子传入的是指针，还有一种叫引用类型，和指针的区别是不需要星号和&，对他的修改会直接改动到原有变量的值。
	//ps:go语言中只有三种引用类型，slice(切片)、map(字典)、channel(管道)

	//函数作为参数
	//在设计模式里，这种方式叫装饰器模式（Decorator Pattern）:允许向一个现有的对象添加新的功能，同时又不改变其结构。
	functionValue(1, 1, add)
	functionValue(1, 1, sub)

	//匿名函数
	functionValue(1, 1, func(i1 int, i2 int) int {
		return i1 * i2
	})

}
