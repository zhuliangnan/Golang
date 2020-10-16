package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
)

/*
go-flags相比标准库flag支持更丰富的数据类型：

所有的基本类型（包括有符号整数int/int8/int16/int32/int64，无符号整数uint/uint8/uint16/uint32/uint64，浮点数float32/float64，布尔类型bool和字符串string）和它们的切片；
map 类型。只支持键为string，值为基础类型的 map；
函数类型。
*/
type Option struct {
	//required非空时，表示对应的选项必须设置值，否则解析时返回ErrRequired错误
	//default用于设置选项的默认值。如果已经设置了默认值，那么required是否设置并不影响，也就是说命令行参数中该选项可以没有
	//示例如下
	//Required    string  `short:"r" long:"required" required:"true"`
	//Default     string  `short:"d" long:"default" default:"default"`
	IntFlag        int          `short:"i" long:"int" default:"1" description:"int flag value"`
	IntSlice       []int        `long:"intslice" description:"int slice flag value"`
	BoolFlag       bool         `long:"bool" description:"bool flag value"`
	BoolSlice      []bool       `long:"boolslice" description:"bool slice flag value"`
	FloatFlag      float64      `long:"float" description:"float64 flag value"`
	FloatSlice     []float64    `long:"floatslice" description:"float64 slice flag value"`
	StringFlag     string       `short:"s" long:"string" description:"string flag value"`
	StringSlice    []string     `long:"strslice" description:"string slice flag value"`
	PtrStringSlice []*string    `long:"pstrslice" description:"slice of pointer of string flag value"`
	Call           func(string) `long:"call" description:"callback"`
	//go-flags还支持 map 类型。虽然限制键必须是string类型，值必须是基本类型，也能实现比较灵活的配置。
	//map类型的选项值中键-值通过:分隔，如key:value，可设置多个。运行代码，传入--intmap选项：
	IntMap map[string]int `long:"intmap" description:"A map from string to int"`
}

func main() {
	var opt Option
	//该函数的唯一要求是有一个字符串类型的参数。解析中每次遇到该选项就会以选项值为参数调用这个函数。
	opt.Call = func(value string) {
		fmt.Println("in callback: ", value)
	}

	_, err := flags.Parse(&opt)
	if err != nil {
		fmt.Println("Parse error:", err)
		return
	}

	fmt.Printf("int flag: %v\n", opt.IntFlag)
	fmt.Printf("int slice flag: %v\n", opt.IntSlice)
	fmt.Printf("bool flag: %v\n", opt.BoolFlag)
	fmt.Printf("bool slice flag: %v\n", opt.BoolSlice)
	fmt.Printf("float flag: %v\n", opt.FloatFlag)
	fmt.Printf("float slice flag: %v\n", opt.FloatSlice)
	fmt.Printf("string flag: %v\n", opt.StringFlag)
	fmt.Printf("string slice flag: %v\n", opt.StringSlice)
	fmt.Println("slice of pointer of string flag: ")
	for i := 0; i < len(opt.PtrStringSlice); i++ {
		fmt.Printf("\t%d: %v\n", i, *opt.PtrStringSlice[i])
	}
	fmt.Printf("int map: %v\n", opt.IntMap)
}
