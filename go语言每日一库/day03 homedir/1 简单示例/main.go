package main

import (
	"fmt"
	"log"
	"os/user"
)

//go-homedir用来获取用户的主目录。 实际上，使用标准库os/user我们也可以得到这个信息：
//在 Darwin 系统上，标准库os/user的使用需要 cgo。所以，任何使用os/user的代码都不能交叉编译。
//但是，大多数人使用os/user的目的仅仅只是想获取主目录。因此，go-homedir库出现了。

func main() {
	u, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Home dir:", u.HomeDir)
}
