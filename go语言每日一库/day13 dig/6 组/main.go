package main

/*
组可以将相同类型的对象放到一个切片中，可以直接使用这个切片。组的定义与上面名字定义类似。可以通过为Provide提供额外的参数：
container.Provide(NewUser("dj", 18), dig.Group("user"))
container.Provide(NewUser("dj2", 18), dig.Group("user"))
也可以在结果对象中添加结构标签group:"user"。

然后我们定义一个参数对象，通过指定同样的结构标签来使用这个切片：
*/
import (
	"fmt"
	"net/http"

	"go.uber.org/dig"
)

type Handler struct {
	Greeting string
	Path     string
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s from %s", h.Greeting, h.Path)
}

func NewHello1Handler() HandlerResult {
	return HandlerResult{
		Handler: Handler{
			Path:     "/hello1",
			Greeting: "welcome",
		},
	}
}

func NewHello2Handler() HandlerResult {
	return HandlerResult{
		Handler: Handler{
			Path:     "/hello2",
			Greeting: "😄",
		},
	}
}

type HandlerResult struct {
	dig.Out

	Handler Handler `group:"server"`
}

type HandlerParams struct {
	dig.In

	Handlers []Handler `group:"server"`
}

func RunServer(params HandlerParams) error {
	mux := http.NewServeMux()
	for _, h := range params.Handlers {
		mux.Handle(h.Path, h)
	}

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

func main() {
	container := dig.New()

	container.Provide(NewHello1Handler)
	container.Provide(NewHello2Handler)

	container.Invoke(RunServer)
}
