// scan.go

package main

import (
	"fmt"
)

// 返回2个函数类型的返回值
func test01(base int) (func(int) int, func(int) int) {
	// 定义2个函数，并返回
	// 相加
	add := func(i int) int {
		base += i
		return base
	}
	// 相减
	sub := func(i int) int {
		base -= i
		return base
	}
	// 返回
	return add, sub
}

func main() {
	f1, f2 := test01(10)
	// base一直是没有消
	fmt.Println(f1(1), f2(2))
	// 此时base是9
	fmt.Println(f1(3), f2(4))

	f3, f4 := test01(10)
	fmt.Println(f3(1), f4(2))
}

/*func main() {


	// 创建一个map 指定key为string类型 val为int类型
	counts := make(map[string]int)
	// 从标准输入流中接收输入数据
	input := bufio.NewScanner(os.Stdin)

	fmt.Printf("Please type in something:\n")

	// 逐行扫描
	for input.Scan() {
		line := input.Text()

		// 输入bye时 结束
		if line == "bye" {
			break
		}

		// 更新key对应的val 新key对应的val是默认0值
		counts[line]++
	}

	// 遍历map统计数据
	for line, n := range counts {
		fmt.Printf("%d : %s\n", n, line)
	}
}*/
