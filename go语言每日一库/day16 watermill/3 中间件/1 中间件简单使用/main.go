package main

import (
	"context"
	"fmt"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"log"
	"time"
)

/*
watermill中内置了几个比较常用的中间件：

IgnoreErrors：可以忽略指定的错误；
Throttle：限流，限制单位时间内处理的消息数量；
Poison：将处理失败的消息以另一个主题发布；
Retry：重试，处理失败可以重试；
Timeout：超时，如果消息处理时间超过给定的时间，直接失败。
InstantAck：直接调用消息的Ack()方法，不管后续成功还是失败；
RandomFail：随机抛出错误，测试时使用；
Duplicator：调用两次处理函数，两次返回的消息都重新发布出去，double~
Correlation：处理函数生成的消息都统一设置成原始消息中的correlation id，方便追踪消息来源；
Recoverer：捕获处理函数中的panic，包装成错误返回。
中间件的使用也是比较简单和直接的：调用router.AddMiddleware()

例如，我们想要把处理返回的消息 double 一下：
router.AddMiddleware(middleware.Duplicator)

想重试？可以
router.AddMiddleware(middleware.Retry{
  MaxRetries:      3,  //最大重试次数为 3
  InitialInterval: time.Millisecond * 100,  //重试初始时间间隔为 100ms
  Logger:          logger,
}.Middleware)
*/

var (
	logger = watermill.NewStdLogger(false, false)
)

func main() {
	router, err := message.NewRouter(message.RouterConfig{}, logger)
	if err != nil {
		panic(err)
	}

	pubSub := gochannel.NewGoChannel(gochannel.Config{}, logger)
	go publishMessages(pubSub)

	router.AddMiddleware(middleware.Duplicator)

	router.AddHandler("myhandler", "in_topic", pubSub, "out_topic", pubSub, myHandler{}.Handler)
	router.AddNoPublisherHandler("print_in_messages", "in_topic", pubSub, printMessages)
	router.AddNoPublisherHandler("print_out_messages", "out_topic", pubSub, printMessages)

	ctx := context.Background()
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
