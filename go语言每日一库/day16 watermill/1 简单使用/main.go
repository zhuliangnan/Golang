package main

/*
watermill是 Go 语言的一个异步消息解决方案，它支持消息重传、保存消息，后启动的订阅者也能收到前面发布的消息。
watermill内置了多种订阅-发布实现，包括Kafka/RabbitMQ，甚至还支持HTTP/MySQL binlog。
当然也可以编写自己的订阅-发布实现。此外，它还提供了监控、限流等中间件。
*/

import (
	"context"
	"log"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
)

func main() {
	//首先，我们创建一个GoChannel对象，它是一个消息管理器。可以调用其Subscribe订阅某个主题（topic）的消息，调用其Publish()以某个主题发布消息
	pubSub := gochannel.NewGoChannel(
		gochannel.Config{},
		watermill.NewStdLogger(false, false),
	)

	//Subscribe()方法会返回一个<-chan *message.Message，一旦该主题有消息发布，GoChannel就会将消息发送到该管道中。订阅者只需监听此管道，接收消息进行处理
	//message.Message这个结构是watermill库的核心，每个消息都会封装到该结构中发送。
	messages, err := pubSub.Subscribe(context.Background(), "example.topic")
	if err != nil {
		panic(err)
	}

	go process(messages)

	publishMessages(pubSub)
}

func process(messages <-chan *message.Message) {
	for msg := range messages {
		log.Printf("received message: %s, payload: %s", msg.UUID, string(msg.Payload))
		//收到的每个消息都需要调用Message的Ack() 方法确认，否则GoChannel会重发当前消息；
		msg.Ack()
	}
}

func publishMessages(publisher message.Publisher) {
	for {
		//Message有一个UUID字段，建议设置为唯一的，方便定位问题。watermill提供方法NewUUID()生成唯一 id。
		msg := message.NewMessage(watermill.NewUUID(), []byte("Hello, world!"))

		if err := publisher.Publish("example.topic", msg); err != nil {
			panic(err)
		}

		time.Sleep(time.Second)
	}
}
