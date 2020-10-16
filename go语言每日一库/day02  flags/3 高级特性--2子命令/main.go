package main

import (
	"errors"
	"fmt"
	"github.com/jessevdk/go-flags"
	"log"
	"strconv"
	"strings"
)

/*
o-flags支持子命令。我们经常使用的 Go 和 Git 命令行程序就有大量的子命令。
例如go version、go build、go run、git status、git commit这些命令中version/build/run/status/commit就是子命令。
使用go-flags定义子命令比较简单：
*/
type MathCommand struct {
	Op     string `long:"op" description:"operation to execute"`
	Args   []string
	Result int64
}

//子命令必须实现go-flags定义的Commander接口：
/*
type Commander interface {
    Execute(args []string) error
}
*/
func (this *MathCommand) Execute(args []string) error {
	if this.Op != "+" && this.Op != "-" && this.Op != "x" && this.Op != "/" {
		return errors.New("invalid op")
	}

	for _, arg := range args {
		num, err := strconv.ParseInt(arg, 10, 64)
		if err != nil {
			return err
		}

		this.Result += num
	}

	this.Args = args
	return nil
}

type Option struct {
	Math MathCommand `command:"math"`
}

func main() {
	var opt Option
	_, err := flags.Parse(&opt)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The result of %s is %d", strings.Join(opt.Math.Args, opt.Math.Op), opt.Math.Result)
}
