package main

/*
viper 可以监听文件修改进而自动重新加载。 其内部使用的就是fsnotify这个库，它是跨平台的
*/
import (
	"log"

	"github.com/fsnotify/fsnotify"
)

func main() {
	/*
		type Watcher struct {
			Events   chan Event
			Errors   chan error
			isClosed bool           // Set to true when Close() is first called
			mu       sync.Mutex     // Map access
			port     syscall.Handle // Handle to completion port
			watches  watchMap       // Map of watches (key: i-number)
			input    chan *input    // Inputs to the reader are sent on this channel
			quit     chan chan<- error
		}
		type Event struct {
			Name string // Name表示发生变化的文件或目录名
			Op   Op     // Op表示具体的变化
		}
		Op取值
		type Op uint32
		const (
		  Create Op = 1 << iota
		  Write
		  Remove
		  Rename
		  Chmod
		)
	*/
	//1. 先调用NewWatcher创建一个监听器； chan
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal("NewWatcher failed: ", err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		defer close(done)

		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Printf("%s %s\n", event.Name, event.Op)
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	//2. 然后调用监听器的Add增加监听的文件或目录
	err = watcher.Add("go语言每日一库/day07 fsnotify/1 简单使用")
	if err != nil {
		log.Fatal("Add failed:", err)
	}
	<-done
}
