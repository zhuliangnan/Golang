package main

import (
	"fmt"
	"github.com/jordan-wright/email"
	"log"
	"net/smtp"
	"os"
	"sync"
	"time"
)

/*
实际上每次调用Send时都会和 SMTP 服务器建立一次连接，如果发送邮件很多很频繁的话可能会有性能问题。email提供了连接池，可以复用网络连接：
*/

func main() {
	ch := make(chan *email.Email, 10)
	p, err := email.NewPool(
		"smtp.163.com:25",
		4,
		smtp.PlainAuth("", "zhuliangnan7410@163.com", "JUBMSUDZGUVSHKYF", "smtp.163.com"),
	)

	if err != nil {
		log.Fatal("failed to create pool:", err)
	}

	var wg sync.WaitGroup
	wg.Add(4)
	for i := 0; i < 4; i++ {
		go func() {
			defer wg.Done()
			for e := range ch {
				err := p.Send(e, 10*time.Second)
				if err != nil {
					fmt.Fprintf(os.Stderr, "email:%v sent error:%v\n", e, err)
				}
			}
		}()
	}

	for i := 0; i < 10; i++ {
		e := email.NewEmail()
		e.From = "zhuliangnan <zhuliangnan7410@163.com>"
		e.To = []string{"2569471714@qq.com"}
		e.Subject = "go语言邮件服务连接池测试"
		e.Text = []byte(fmt.Sprintf("Awesome Web %d", i+1))
		ch <- e
	}

	close(ch)
	wg.Wait()
}
