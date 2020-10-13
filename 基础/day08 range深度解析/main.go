package main

import (
	"fmt"
	"sort"
)

func main() {

	//切片迭代
	fmt.Println("----切片迭代----")
	nums := []int{1, 2, 3, 4, 5, 6}
	for k, v := range nums {
		nums[k] = 9
		fmt.Printf("key : %v  value:  %v  \n", k, v)
	}
	fmt.Println("----切片迭代----")
	for k1, v1 := range nums {
		fmt.Printf("key : %v  value:  %v  \n", k1, v1)
	}
	//这和迭代方式非常适合用for-range语句，如果减少赋值，直接索引num[key]可以减少损耗。如下
	// for k, _:= range nums

	fmt.Println("----map迭代----")
	//map 迭代
	//注意，从 Go1开始，遍历的起始节点就是随机了。  参考 java 里面的 map
	kvs := map[string]string{
		"name": "zhu",
		"sex":  "男",
	}

	for mk, mv := range kvs {
		fmt.Printf(" mk : %v  mv : %v \n", mk, mv)
	}

	//在以前的函数中for-range语句中只获取 key 值，然后跟据 key 值获取 value 值，虽然看似减少了一次赋值，但通过 key 值查找 value 值的性能消耗可能高于赋值消耗。
	//所以能否优化取决于 map 所存储数据结构特征、结合实际情况进行。

	fmt.Println("----字符串迭代----")
	for sk, sv := range "hello world" {
		//注意这里单个字符输出的是ASCII码，
		//用 %c 代表输出字符
		fmt.Printf(" sk : %v  sv : %c  \n", sk, sv)
	}

	fmt.Println("----channel迭代----")
	//channel （如果不会可以先 mark 下，详细参考后续：go 的并发特性章节）
	ch := make(chan int, 10)
	ch <- 11
	ch <- 12

	close(ch) // 不用的时候记得关掉,不关掉又没有另一个goroutine存在会死锁哦，可以注释掉这一句体验死锁

	for x := range ch {
		fmt.Println(x)
	}

	fmt.Println("----结构体迭代----")
	people := []struct {
		string
		int
	}{
		{"zhu", 20},
		{"nan", 30},
	}

	for tk, tv := range people {
		fmt.Printf("tk : %v   tv : %v  \n", tk, tv)
	}
	//注意：由于循环开始前循环次数就已经确定了，所以循环过程中新添加的元素是没办法遍历到的。
	//由于range遍历时value是值的 **拷贝**，如果这个时候遍历声明的结构体时，修改value，原结构体不会发生任何变化！
	for _, v := range people {
		v.string = "li"
		v.int = 30
		//你这里修改的 只是拷贝的值
		fmt.Printf("打印修改拷贝的值 v : %v  \n", v)
	}
	//在这里可以看到 值没有被修改
	for _, v := range people {
		fmt.Printf(" v : %v  \n", v)
	}

	//追加操作
	var runes []rune
	for _, r := range "hello , 世界" {
		runes = append(runes, r)
	}
	fmt.Printf("%q  长度: %v \n ", runes, len(runes))

	var y []int
	for i := 0; i < 10; i++ {
		y = append(y, i)
		fmt.Printf("%d cap=%d \t %v \n", i, cap(y), y)

	}

	var x []int
	x = append(x, 1)
	x = append(x, 2, 3)
	x = append(x, 4, 5, 6)
	x = append(x, x...) //追加一个 slice x  我加我自己
	fmt.Println(x)

	var k []int
	k = appendInt(k, 1, 2, 3)
	k = appendInt(k, 4, 5)
	k = appendInt(k, k...)
	fmt.Println(k)

	fmt.Println("----------")
	var a []int
	b := make([]int, 2, 10)
	a = append(a, 1, 2, 3)
	fmt.Println(a)
	copy(b, a)
	fmt.Println(b)

	slice2 := []int{0, 3, 0, 1, 2, 0}
	sort.Ints(slice2)
	fmt.Println(slice2)

}

// 反转int
func reverse(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

// 反转字符串
func reverseString(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}

//x []int--- 要追加的切片 y--要追加的值
func appendInt(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)
	if zlen <= cap(x) {
		// There is	room to	grow. Extend the slice
		z = x[:zlen]
	} else {
		// There is insufficient space .Allocate a new array
		// Grow	by doubling,for amortized linear complexity
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x) // copy函数的第一个参数是要复制的目标slice，第二个参数是源slice
	}
	copy(z[len(x):], y)
	return z

}
