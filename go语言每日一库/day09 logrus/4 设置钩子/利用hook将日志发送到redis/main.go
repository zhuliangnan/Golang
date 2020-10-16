package main

/*
mgorus：将日志发送到 mongodb；
logrus-redis-hook：将日志发送到 redis；
logrus-amqp：将日志发送到 ActiveMQ。
*/

import (
	"io/ioutil"

	logredis "github.com/rogierlommers/logrus-redis-hook"
	"github.com/sirupsen/logrus"
)

func init() {
	hookConfig := logredis.HookConfig{
		Host:     "192.168.248.128",
		Key:      "mykey",
		Format:   "v0",
		App:      "aweosome",
		Port:     6379,
		Hostname: "my_app_hostname",
		DB:       0, // optional
		TTL:      3600,
	}

	hook, err := logredis.NewHook(hookConfig)
	if err == nil {
		logrus.AddHook(hook)
	} else {
		logrus.Errorf("logredis error: %q", err)
	}
}

func main() {
	logrus.Info("just some info logging...")

	logrus.WithFields(logrus.Fields{
		"animal": "walrus",
		"foo":    "bar",
		"this":   "that",
	}).Info("additional fields are being logged as well")

	logrus.SetOutput(ioutil.Discard)
	logrus.Info("This will only be sent to Redis")
}
