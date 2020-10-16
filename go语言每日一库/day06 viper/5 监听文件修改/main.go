package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"log"
	"time"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("go语言每日一库/day06 viper/4 保存配置/config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("read config failed: %v", err)
	}

	//只需要调用viper.WatchConfig，viper 会自动监听配置修改。如果有修改，重新加载的配置。
	viper.WatchConfig()

	//这样文件修改时会执行这个回调。
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("Config file:%s Op:%s\n", e.Name, e.Op)
	})

	fmt.Println("redis port before sleep: ", viper.Get("redis.port"))
	time.Sleep(time.Second * 10)
	fmt.Println("redis port after sleep: ", viper.Get("redis.port"))
}
