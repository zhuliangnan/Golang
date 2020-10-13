package main

import "fmt"

func main() {

	nums := []int{1, 2, 3, 4, 5, 6}

	//for 循环方式 1 和 c++、java 相似
	for i := 0; i < len(nums); i++ {
		fmt.Println(i)
	}

	fmt.Println("----------------")
	//for 循环方式 2 省略赋值和++
	a, b := 1, 5
	for a < b {
		fmt.Println(a)
		a++
	}
	fmt.Println("----------------")
	//for 循环方式 3 迭代
	//优点：不用引入无意义的变量
	//缺点：不是直接索引，如果数据量极大会有性能损耗

	for index, value := range nums {
		fmt.Printf("key: %v  value: %v  \n", index, value)

	}
	//当然，你可以把方式 3 中 index 去掉,用_忽略掉key 或者 用_忽略掉value
	for _, value := range nums {
		fmt.Printf(" value: %v  \n", value)

	}
	//如果你想忽略掉 value，直接用 key也是可以的，这样就消除了迭代方式的缺点！
	//当然 这样只默认输出第一个也就是key 所以取value还是需要用 nums[index]
	for index := range nums {
		fmt.Printf("value: %v  \n", nums[index])

	}
	//死循环
	//go 里面没有while 要无限循环的话 可以使用
	//这里为了演示加入了break
	for {
		break
	}
	fmt.Println("------break、continue 测试----------")
	//break、continue
	i := 0
	for {
		i++
		if i >= 5 {
			fmt.Println("跳出循环")
			break
		}
		if i == 2 {
			fmt.Println("触发continue ")
			continue
		}

		fmt.Printf("输出 i : %v  \n", i)

	}

}
