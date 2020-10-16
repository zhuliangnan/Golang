package main

import (
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
	"time"
)

func main() {
	logrus.SetLevel(logrus.TraceLevel)
	/*
		logrus支持两种日志格式，文本和 JSON，默认为文本格式。可以通过logrus.SetFormatter设置日志格式：
	*/
	//logrus.SetFormatter(&logrus.JSONFormatter{})

	//除了内置的TextFormatter和JSONFormatter，还有不少第三方格式支持。我们这里介绍一个nested-logrus-formatter。
	logrus.SetFormatter(&nested.Formatter{
		//以通过设置HideKeys为true隐藏键，只输出值；
		//HideKeys:    true,
		TimestampFormat: time.RFC3339,
		//可以设置FieldsOrder定义输出字段顺序
		FieldsOrder: []string{"component", "category"},
	})

	logrus.WithFields(logrus.Fields{
		"name": "dj",
		"age":  18,
	}).Info("info msg")

	logrus.Trace("trace msg")
	logrus.Debug("debug msg")
	logrus.Info("info msg")
	logrus.Warn("warn msg")
	logrus.Error("error msg")
	logrus.Fatal("fatal msg")
	logrus.Panic("panic msg")
}
