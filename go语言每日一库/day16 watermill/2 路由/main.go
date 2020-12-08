package main

import (
	"context"
	"fmt"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"log"
	"time"
)

/*
在实际应用中，我们通常想要监控、重试、统计等一些功能。而且上面的例子中，每个消息处理结束需要手动调用Ack()方法，消息管理器才会下发后面一条信息，很容易遗忘。
还有些时候，我们有这样的需求，处理完某个消息后，重新发布另外一些消息。

这些功能都是比较通用的，为此watermill提供了路由（Router）功能。直接拿来官网的图：

路由其实管理多个订阅者，每个订阅者在一个独立的goroutine中运行，彼此互不干扰。
订阅者收到消息后，交由注册时指定的处理函数（HandlerFunc）。
路由还可以设置插件（plugin）和中间件（middleware），插件是定制路由的行为，而中间件是定制处理器的行为。
处理器处理消息后会返回若干消息，这些消息会被路由重新发布到（另一个）管理器中。

使用路由还有个好处，处理器返回时，若无错误，路由会自动调用消息的Ack()方法；若发生错误，路由会调用消息的Nack()方法通知管理器重发这条消息。
*/
var (
	logger = watermill.NewStdLogger(false, false)
)

func main() {
	//我们创建一个路由：
	router, err := message.NewRouter(message.RouterConfig{}, logger)
	if err != nil {
		panic(err)
	}

	//我们创建一个GoChannel对象，它是一个消息管理器。可以调用其Subscribe订阅某个主题（topic）的消息，调用其Publish()以某个主题发布消息
	pubSub := gochannel.NewGoChannel(gochannel.Config{}, logger)
	go publishMessages(pubSub)

	//然后为路由注册处理器。注册的处理器有两种类型，一种是： router.AddHandler
	//该方法的作用是创建一个名为handlerName的处理器，监听subscriber中主题为subscribeTopic的消息，收到消息后调用handlerFunc处理，将返回的消息以主题publishTopic发布到publisher中。
	router.AddHandler("myhandler", "in_topic", pubSub, "out_topic", pubSub, myHandler{}.Handler)
	//另外一种处理器是下面这种形式：这种形式的处理器只处理接收到的消息，不发布新消息。
	router.AddNoPublisherHandler("print_in_messages", "in_topic", pubSub, printMessages)
	router.AddNoPublisherHandler("print_out_messages", "out_topic", pubSub, printMessages)

	ctx := context.Background()
	//调用router.Run()运行这个路由。
	if err := router.Run(ctx); err != nil {
		panic(err)
	}
}

func publishMessages(publisher message.Publisher) {
	for {
		msg := message.NewMessage(watermill.NewUUID(), []byte("Hello, world!"))
		if err := publisher.Publish("in_topic", msg); err != nil {
			panic(err)
		}

		time.Sleep(time.Second)
	}
}

func printMessages(msg *message.Message) error {
	fmt.Printf("\n> Received message: %s\n> %s\n>\n", msg.UUID, string(msg.Payload))
	return nil
}

type myHandler struct {
}

func (m myHandler) Handler(msg *message.Message) ([]*message.Message, error) {
	log.Println("myHandler received message", msg.UUID)

	msg = message.NewMessage(watermill.NewUUID(), []byte("message produced by myHandler"))
	return message.Messages{msg}, nil
}
