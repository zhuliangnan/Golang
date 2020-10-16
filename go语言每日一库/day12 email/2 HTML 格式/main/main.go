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
	e.Subject = "go语言邮件服务测试"
	e.HTML = []byte(`
  <ul>
<li><a "https://darjun.github.io/2020/01/10/godailylib/flag/">Go 每日一库之 flag</a></li>
<li><a "https://darjun.github.io/2020/01/10/godailylib/go-flags/">Go 每日一库之 go-flags</a></li>
<li><a "https://darjun.github.io/2020/01/14/godailylib/go-homedir/">Go 每日一库之 go-homedir</a></li>
<li><a "https://darjun.github.io/2020/01/15/godailylib/go-ini/">Go 每日一库之 go-ini</a></li>
<li><a "https://darjun.github.io/2020/01/17/godailylib/cobra/">Go 每日一库之 cobra</a></li>
<li><a "https://darjun.github.io/2020/01/18/godailylib/viper/">Go 每日一库之 viper</a></li>
<li><a "https://darjun.github.io/2020/01/19/godailylib/fsnotify/">Go 每日一库之 fsnotify</a></li>
<li><a "https://darjun.github.io/2020/01/20/godailylib/cast/">Go 每日一库之 cast</a></li>
</ul>
  `)
	err := e.Send("smtp.163.com:25", smtp.PlainAuth("", "zhuliangnan7410@163.com", "JUBMSUDZGUVSHKYF", "smtp.163.com"))
	if err != nil {
		log.Fatal(err)
	}
}
