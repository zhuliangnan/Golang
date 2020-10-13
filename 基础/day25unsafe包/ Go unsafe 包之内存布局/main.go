package main

import (
	"fmt"
	"unsafe"
)

func main() {


	fmt.Println("-------Sizeof函数-------")
	//Sizeof函数可以返回一个类型所占用的内存大小，这个大小只有类型有关，
	//和类型对应的变量存储的内容大小无关，比如bool型占用一个字节、int8也占用一个字节。
	fmt.Println(unsafe.Sizeof(true))
	fmt.Println(unsafe.Sizeof(int8(0)))
	fmt.Println(unsafe.Sizeof(int16(10)))
	fmt.Println(unsafe.Sizeof(int32(10000000)))
	fmt.Println(unsafe.Sizeof(int64(10000000000000)))
	fmt.Println(unsafe.Sizeof(int(10000000000000000)))

	fmt.Println("-------Alignof函数-------")
	//Alignof返回一个类型的对齐值，也可以叫做对齐系数或者对齐倍数。
	//对齐值是一个和内存对齐有关的值，合理的内存对齐可以提高内存读写的性能，
	//关于内存对齐的知识可以参考相关文档，这里不展开介绍。
	var b bool
	var i8 int8
	var i16 int16
	var i64 int64

	var f32 float32

	var s string

	var m map[string]string

	var p *int32

	fmt.Println(unsafe.Alignof(b))
	fmt.Println(unsafe.Alignof(i8))
	fmt.Println(unsafe.Alignof(i16))
	fmt.Println(unsafe.Alignof(i64))
	fmt.Println(unsafe.Alignof(f32))
	fmt.Println(unsafe.Alignof(s))
	fmt.Println(unsafe.Alignof(m))
	fmt.Println(unsafe.Alignof(p))

	fmt.Println("-------Offsetof 函数-------")
	//Offsetof函数只适用于struct结构体中的字段相对于结构体的内存位置偏移量。结构体的第一个字段的偏移量都是0.
	//根据字段的偏移量，我们可以定位结构体的字段，进而可以读写该结构体的字段，哪怕他们是私有的
	var u1 user1

	fmt.Println(unsafe.Offsetof(u1.b))
	fmt.Println(unsafe.Offsetof(u1.i))
	fmt.Println(unsafe.Offsetof(u1.j))

	fmt.Println("-------有意思的struct大小-------")

	var u2 user2
	var u3 user3
	var u4 user4
	var u5 user5
	var u6 user6
//因为有内存对齐存在，编译器使用了内存对齐，那么最后的大小结果就不一样了
	fmt.Println("u1 size is ",unsafe.Sizeof(u1))
	fmt.Println("u2 size is ",unsafe.Sizeof(u2))
	fmt.Println("u3 size is ",unsafe.Sizeof(u3))
	fmt.Println("u4 size is ",unsafe.Sizeof(u4))
	fmt.Println("u5 size is ",unsafe.Sizeof(u5))
	fmt.Println("u6 size is ",unsafe.Sizeof(u6))

}
type user1 struct {
	b byte
	i int32
	j int64
}

type user2 struct {
	b byte
	j int64
	i int32
}

type user3 struct {
	i int32
	b byte
	j int64
}

type user4 struct {
	i int32
	j int64
	b byte
}

type user5 struct {
	j int64
	b byte
	i int32
}

type user6 struct {
	j int64
	i int32
	b byte
}