package main

import "fmt"

func main() {

	fmt.Println("------------数组----------------")
	//数组中是固定长度的连续空间（内存区域）
	//数组中所有元素的类型是一样的
	var array1 [10]int             //1.常规的数组声明方法,它的类型是[10]int,数组长度被认为是数组类型的一部分
	var array2 = [10]int{1, 2, 3}  //2.常用用法,给定初始化的值,未被初始化的值以该类型的零值填充
	array3 := [...]int{1, 2, 3, 4} //3.不给定长度,使用...关键词告知编译期根据初始化的元素个数推断长度
	fmt.Printf("array1 type:%T len:%d, cap:%d\n", array1, len(array1), cap(array1))
	fmt.Printf("array2 type:%T len:%d, cap:%d\n", array2, len(array2), cap(array2))
	fmt.Printf("array2 type:%T len:%d, cap:%d\n", array3, len(array3), cap(array3))

	//大家可以看到 array2的长度和容量也是10 为什么呢？ 这个是因为未被初始化的值以该类型的零值填充
	//下面我们输出一下array2
	for k, v := range array2 {
		fmt.Printf("key : %v  value:  %v  \n", k, v)
	}

	//Tip
	//在Go中数组的长度属于数组类型的一部分
	//所以在函数调用时如果参数类型是数组，那么每次传参都会发生一次数组的copy，这个对性能的影响还是比较大的，所以我们一般在Go中都是使用slice这种数据结构。
	//在Go中数组是值类型 这点要和C/C++区分

	//多维数组
	//声明二维数组，只要 任意加中括号，可以声明更多维，相应占用空间指数上指
	var arr [3][3]int
	//赋值
	arr = [3][3]int{
		{1, 2, 3},
		{2, 3, 4},
		{3, 4, 5},
	}
	fmt.Println(arr)

	fmt.Println("------------切片----------------")

	//切片
	//一种可以动态增长的数组。

	//常用的切片构造方式
	var slen int = 10
	var scap int = 20

	var slice1 = make([]int, slen, scap) //这种方法可读性比较好,显示指定了slice的长度和容量
	slice2 := []int{1, 2, 3, 4}
	slice3 := []int{} //当然生成一个空的切片用这个就行 var slice3 []int
	//切片的底层是一个数组，除此之外Go的切片和数组再无联系。
	//切片的长度和容量很可能是不一样的，但是容量一定大于等于长度，可以使用len和cap方法查看：
	fmt.Printf("slice1 type:%T len:%d, cap:%d\n", slice1, len(slice1), cap(slice1))
	fmt.Printf("slice2 type:%T len:%d, cap:%d\n", slice2, len(slice2), cap(slice2))
	fmt.Printf("slice3 type:%T len:%d, cap:%d\n", slice3, len(slice3), cap(slice3))
	//从输出我们可以看出，切片的长度是不属于切片类型的一部分的（显然，因为切片的长度是变化的）

	//为了证明切片是引用类型 下面我们截取(前闭后开) 一部分用来测试
	tmp := slice2[1:3] // 截取下标  1  2
	tmp[0] += 2

	fmt.Println(slice2)
	fmt.Println(tmp)
	//下面看一下 截取后的 长度和容量
	fmt.Printf("tmp type:%T len:%d, cap:%d\n", tmp, len(tmp), cap(tmp))

	//案例
	months := [...]string{1: "January", 2: "February", 3: "March", 4: "April", 5: "May", 6: "June", 7: "July", 8: "August", 9: "September", 10: "October", 11: "November", 12: "December"}
	Q2 := months[4:7]
	summer := months[6:9]
	fmt.Println(Q2)     // ["April" "May" "June"]
	fmt.Println(summer) // ["June" "July" "August"]

	//fmt.Println(summer[:20]) //slice bounds out of range [:20] with capacity 7
	endlessSummer := summer[:5]
	fmt.Println(endlessSummer) //[June July August September October]
	//当然这种写法也是不规范的
	if summer == nil {
		fmt.Println("切片为nil")
	}

}

func equal(x, y []string) bool {

	if len(x) != len(y) {
		return false
	}
	for k := range x {
		if x[k] != x[k] {
			return false
		}

	}
	return true
}
