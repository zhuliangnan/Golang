package day22性能测试__基准测试

import (
	"fmt"
	"strconv"
	"testing"
)

/*
基准测试的代码文件必须以_test.go结尾
基准测试的函数必须以Benchmark开头，必须是可导出的
基准测试函数必须接受一个指向Benchmark类型的指针作为唯一参数
基准测试函数不能有返回值
b.ResetTimer是重置计时器，这样可以避免for循环之前的初始化代码的干扰
最后的for循环很重要，被测试的代码要放到循环里
b.N是基准测试框架提供的，表示循环的次数，因为需要反复调用测试的代码，才可以评估性能
*/
//int类型转为string类型
func BenchmarkSprintf(b *testing.B) {

	num := 10
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fmt.Printf("%d", num)

	}

}
func BenchmarkFormat(b *testing.B) {
	num := int64(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		strconv.FormatInt(num, 10)
	}
}

func BenchmarkItoa(b *testing.B) {
	num := 10
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		strconv.Itoa(num)
	}
}
