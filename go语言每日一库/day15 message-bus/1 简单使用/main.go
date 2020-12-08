package main

/*
在一个涉及多模块交互的系统中，如果模块的交互需要手动去调用对方的方法，
那么代码的耦合度就太高了。所以产生了异步消息通信。
实际上，各种各样的消息队列都是基于异步消息的。
不过它们大部分都有着非常复杂的设计，很多被设计成一个独立的软件来使用。
今天我们介绍一个非常小巧的异步消息通信库[message-bus]
它只能在一个进程中使用。源代码只有一个文件
*/

import (
	"fmt"
	"sync"

	messagebus "github.com/vardius/message-bus"
)

func main() {
	queueSize := 100
	bus := messagebus.New(queueSize)

	var wg sync.WaitGroup
	wg.Add(2)

	//message-bus承担了模块间消息分发的角色。模块 A 和 模块 B 先向message-bus订阅主题（topic）
	_ = bus.Subscribe("topic", func(v bool) {
		defer wg.Done()
		fmt.Println(v)
	})

	_ = bus.Subscribe("topic", func(v bool) {
		defer wg.Done()
		fmt.Println(v)
	})

	//调用Publish()向管理器发布主题消息
	bus.Publish("topic", true)
	wg.Wait()
}
