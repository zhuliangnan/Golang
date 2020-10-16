package main

import "github.com/sirupsen/logrus"

func main() {

	//为了能看到Trace和Debug日志，我们在main函数第一行设置日志级别为TraceLevel。
	logrus.SetLevel(logrus.TraceLevel)

	//调用logrus.SetReportCaller(true)设置在输出日志中添加文件名和方法信息：
	logrus.SetReportCaller(true)

	//有时候需要在输出中添加一些字段，可以通过调用logrus.WithField和logrus.WithFields实现
	//logrus.WithFields接受一个logrus.Fields类型的参数，其底层实际上为map[string]interface{}
	//这里再Info输出级别上加了字段
	logrus.WithFields(logrus.Fields{
		"name": "dj",
		"age":  18,
	}).Info("info msg")

	//如果在一个函数中的所有日志都需要添加某些字段，可以使用WithFields的返回值。例如在 Web 请求的处理器中，日志都要加上user_id和ip字段：
	requestLogger := logrus.WithFields(logrus.Fields{
		"user_id": 10010,
		"ip":      "192.168.32.15",
	})
	requestLogger.Info("info msg")
	requestLogger.Error("error msg")

	logrus.Trace("trace msg")
	logrus.Debug("debug msg")
	logrus.Info("info msg") //默认
	logrus.Warn("warn msg")
	logrus.Error("error msg")
	//由于logrus.Fatal会导致程序退出，下面的logrus.Panic不会执行到。
	logrus.Fatal("fatal msg")
	logrus.Panic("panic msg")

}
