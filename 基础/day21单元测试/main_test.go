package main

//含有单元测试代码的go文件必须以_test.go结尾，Go语言测试工具只认符合这个规则的文件
import (
	"mytest"
	"testing"
)

/*func TestAdd(t *testing.T) {
	sum := mytestadd.Add(1,2)
	if sum == 3 {
		t.Log("the result is ok")
	} else {
		t.Fatal("the result is wrong")
	}
}*/

//表组测试---有好几个不同的输入以及输出组成
func TestAdd(t *testing.T) {
	sum := mytest.Add(1, 2)
	if sum == 3 {
		t.Log("the result is ok")
	} else {
		t.Fatal("the result is wrong")
	}

	sum = mytest.Add(3, 4)
	if sum == 7 {
		t.Log("the result is ok")
	} else {
		t.Fatal("the result is wrong")
	}

}
