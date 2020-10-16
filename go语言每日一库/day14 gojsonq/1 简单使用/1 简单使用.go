package main

import (
	"fmt"
	"github.com/thedevsaddam/gojsonq"
)

func main() {
	content := `{
  "user": {
    "name": "dj",
    "age": 18,
    "address": {
      "provice": "shanghai",
      "district": "xuhui"
    },
    "hobbies":["chess", "programming", "game"]
  }
}`

	//首先调用gojsonq.New()创建一个JSONQ的对象
	gq := gojsonq.New().FromString(content)
	district := gq.Find("user.address.district")
	fmt.Println(district)
	//在查询之后，我们手动调用了一次Reset()方法。因为JSONQ对象在调用Find方法时，
	//内部会记录当前的节点，下一个查询会从上次查找的节点开始。
	//也就是说如果我们注释掉jq.Reset()，第二个Find()方法实际上查找的是user.address.district.user.hobbies.[0]，自然就返回nil了
	/*
		除此之外，gojsonq也提供了另外一种方式。如果你想要保存当前查询的一些状态信息，
		可以调用JSONQ的Copy方法返回一个初始状态下的对象，它们会共用底层的 JSON 字符串和解析后的对象。上面的gq.Reset()可以由下面这行代码代替：

		//调用JSONQ的Copy方法返回一个初始状态下的对象
		//gpCopy := gq.Copy()
	*/
	gq.Reset()

	hobby := gq.Find("user.hobbies.[0]")
	fmt.Println(hobby)
}
