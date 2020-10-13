package main

import (
	"fmt"
	"unsafe"
)

func main() {

	i := 1
	ip := &i
	//var fp *float64 = *float64(ip)
	//fmt.Println(fp)

	fmt.Println("-------------unsafe.Pointer----------")
	/*
		任何指针都可以转换为unsafe.Pointer
		unsafe.Pointer可以转换为任何指针
		uintptr可以转换为unsafe.Pointer --*T是不能计算偏移量的，也不能进行计算，但是uintptr可以，所以我们可以把指针转为uintptr再进行偏移计算，这样我们就可以访问特定的内存了，达到对不同的内存读写的目的。
		unsafe.Pointer可以转换为uintptr
	*/
	var fp *float64 = (*float64)(unsafe.Pointer(ip))
	*fp = *fp * 3
	fmt.Println(i)

	u := new(user)
	fmt.Println(*u)

	pName := (*string)(unsafe.Pointer(u)) //第一个修改user的name值的时候，因为name是第一个字段，所以不用偏移，我们获取user的指针，然后通过unsafe.Pointer转为*string进行赋值操作即可。
	*pName = "张三"

	//因为age不是第一个字段，所以我们需要内存偏移，内存偏移牵涉到的计算只能通过uintptr，
	//所我们要先把user的指针地址转为uintptr，然后我们再通过unsafe.Offsetof(u.age)获取需要偏移的值，进行地址运算(+)偏移即可。
	pAge := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(u)) + unsafe.Offsetof(u.age)))
	*pAge = 20

	fmt.Println(*u)
	//逻辑上看，以上代码不会有什么问题，但是这里会牵涉到GC，如果我们的这些临时变量被GC，那么导致的内存操作就错了，我们最终操作的，就不知道是哪块内存了，会引起莫名其妙的问题。

}

type user struct {
	name string
	age  int
}
