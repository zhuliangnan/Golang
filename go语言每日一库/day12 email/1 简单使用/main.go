package main

import (
	"github.com/jordan-wright/email"
	"log"
	"net/smtp"
)

func main() {
	//先调用NewEmail创建一封邮件；
	e := email.NewEmail()
	/*
		设置From发送方，To接收者，Subject邮件主题（标题），Text设置邮件内容；
		然后调用Send发送，参数1是 SMTP 服务器的地址，参数2为验证信息。
	*/
	e.From = "zhuliangnan <zhuliangnan7410@163.com>"
	e.To = []string{"2569471714@qq.com"}
	/*
	       //抄送
	       e.Cc = []string{"test1@126.com", "test2@126.com"}
	       //秘密抄送
	   	e.Bcc = []string{"secret@126.com"}
	*/
	e.Subject = "go语言邮件服务测试"
	e.Text = []byte("Text Body is, of course, supported!")
	err := e.Send("smtp.163.com:25", smtp.PlainAuth("", "zhuliangnan7410@163.com", "JUBMSUDZGUVSHKYF", "smtp.163.com"))
	if err != nil {
		log.Fatal(err)
	}
}
