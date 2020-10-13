//字符串连接
package main

import (
	"fmt"
	"os"
)

func main() {

	//字符串连接的一种方式

	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)

	//如果连接涉及的数据量很大，这种方式代价高昂。一种简单且高效的解决方案是使 用	strings	包的	Join	函数：

	//fmt.Println(strings.Join(os.Args[1:], ""))

}
