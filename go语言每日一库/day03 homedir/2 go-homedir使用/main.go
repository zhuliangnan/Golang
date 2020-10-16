package main

/*
go-homedir有两个功能：

Dir：获取用户主目录；
Expand：将路径中的第一个~扩展成用户主目录。
*/
import (
	"fmt"
	"log"

	"github.com/mitchellh/go-homedir"
)

func main() {

	/*
		由于Dir的调用可能涉及一些系统调用和外部执行命令，多次调用费性能。
		所以go-homedir提供了缓存的功能。默认情况下，缓存是开启的。 我们也可以将DisableCache设置为false来关闭它。
	*/
	//homedir.DisableCache = false

	dir, err := homedir.Dir()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Home dir:", dir)

	dir = "~/golang/src"
	expandedDir, err := homedir.Expand(dir)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Expand of %s is: %s\n", dir, expandedDir)
}
