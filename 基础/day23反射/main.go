package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age  int
}

func main() {

	u := User{"张三", 20}
	t := reflect.TypeOf(u)
	v := reflect.ValueOf(u)
	fmt.Println(t)
	fmt.Println(v)

	fmt.Printf("%T\n", u)
	fmt.Printf("%v\n", u)

	fmt.Println("-------reflect.Value转原始类型--------")

	//reflect.Value又同时持有一个对象的reflect.Value和reflect.Type,所以我们可以通过reflect.Value的Interface方法实现还原
	u1 := v.Interface().(User)
	fmt.Println(u1)
	t1 := v.Type()
	fmt.Println(t1)

	fmt.Println("-------获取类型底层类型--------")

	//获取类型底层类型
	fmt.Println(t.Kind())
	fmt.Println(v.Kind())

	fmt.Println("-------遍历字段和方法--------")
	for i := 0; i < t.NumField(); i++ {
		fmt.Println(t.Field(i).Name)
	}

	for i := 0; i < t.NumMethod(); i++ {
		fmt.Println(t.Method(i).Name)
	}

	//因为reflect.ValueOf函数返回的是一份值的拷贝，所以前提是我们是传入要修改变量的地址。
	//其次需要我们调用Elem方法找到这个指针指向的值。 最后我们就可以使用SetInt方法修改值了。
	fmt.Println("-------修改字段的值--------")
	x := 2
	k := reflect.ValueOf(&x)
	k.Elem().SetInt(100)
	fmt.Println(x)

	fmt.Println("-------动态调用方法--------")

	mPrint := v.MethodByName("Print")
	//获取参数
	args := []reflect.Value{reflect.ValueOf("前缀")}
	fmt.Println(mPrint.Call(args))

}
func (u User) Print(prfix string) {
	fmt.Printf("%s:Name is %s,Age is %d", prfix, u.Name, u.Age)
}
