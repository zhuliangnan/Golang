package main

import (
	"bytes"
	"io"
	"log"
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	writer1 := &bytes.Buffer{}
	writer2 := os.Stdout
	writer3, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		log.Fatalf("create file log.txt failed: %v", err)
	}
	//默认情况下，日志输出到io.Stderr。可以调用logrus.SetOutput传入一个io.Writer参数。
	logrus.SetOutput(io.MultiWriter(writer1, writer2, writer3))
	logrus.Info("info1 msg")
	logrus.Info("info2 msg")
	logrus.Info("info3 msg")
	logrus.Info("info4 msg")
	logrus.Info("info5 msg")
}
