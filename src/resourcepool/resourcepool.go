package resourcepool

import (
	"errors"
	"io"
	"log"
	"sync"
)

//一个安全的资源池，被管理的资源必须都实现io.Close接口
type Pool struct {
	m       sync.Mutex
	res     chan io.Closer            //res字段是一个有缓冲的通道，用来保存共享的资源，这个通道的大小，在初始化Pool的时候就指定的
	factory func() (io.Closer, error) //当需要一个新的资源时，可以通过这个函数创建
	closed  bool
}

var ErrPoolClosed error = errors.New("资源池已经关闭")

//这个函数创建一个资源池，它接收两个参数，一个fn是创建新资源的函数；还有一个size是指定资源池的大小。
func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("size的值太小了")
	}

	return &Pool{
		factory: fn,
		res:     make(chan io.Closer, size),
	}, nil

}

//从资源池里获取一个资源
func (p *Pool) Acquire() (io.Closer, error) {
	select {
	case r, ok := <-p.res:
		log.Println("Acquire:共享资源")
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil
	default:
		log.Println("Acquire:新生成资源")
		return p.factory()
	}

}

//关闭资源池
func (p *Pool) Close() {
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		return
	}
	p.closed = true

	//关闭通道 ，不让写入了
	close(p.res)

	//关闭通道里的资源
	for r := range p.res {
		r.Close()
	}

}

//释放资源池--释放资源本质上就会把资源再发送到缓冲通道中
func (p *Pool) Release(r io.Closer) {
	//保证该操作和Close方法的操作是安全的
	p.m.Lock()
	defer p.m.Unlock()

	//资源池都关闭了，就剩下这一个没有释放的资源了，释放即可
	if p.closed {
		r.Close()
		return
	}

	select {
	case p.res <- r:
		log.Println("资源释放到池子里了")
	default:
		log.Println("资源池满了，释放这个资源吧")
		r.Close()

	}

}
