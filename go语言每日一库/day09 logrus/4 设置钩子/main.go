package main

import "github.com/sirupsen/logrus"

/*
还可以为logrus设置钩子，每条日志输出前都会执行钩子的特定方法。
所以，我们可以添加输出字段、根据级别将日志输出到不同的目的地。
logrus也内置了一个syslog的钩子，将日志输出到syslog中。
这里我们实现一个钩子，在输出的日志中增加一个app=awesome-web字段。

钩子需要实现logrus.Hook接口：
type Hook interface {
  Levels() []Level  //Levels()方法返回感兴趣的日志级别，输出其他日志时不会触发钩子
  Fire(*Entry) error  //Levels()方法返回感兴趣的日志级别
}
*/
type AppHook struct {
	AppName string
}

func (h *AppHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *AppHook) Fire(entry *logrus.Entry) error {
	entry.Data["app"] = h.AppName
	return nil
}

func main() {
	h := &AppHook{AppName: "awesome-web"}
	logrus.AddHook(h)

	logrus.Info("info msg")
	logrus.Warn("warn msg")
}
