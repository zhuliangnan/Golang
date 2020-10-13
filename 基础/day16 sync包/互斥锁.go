package main

import (
	"fmt"
	"sync"
	"time"
)

type Bank struct {
	sync.Mutex
	saving map[string]int
}

func NewBank() *Bank {
	b := &Bank{
		saving: make(map[string]int),
	}
	return b

}

// Deposit 存款
func (b *Bank) Deposit(name string, amount int) {
	b.Lock()
	defer b.Unlock()

	fmt.Println("开始存钱 存", amount, "元")
	if _, ok := b.saving[name]; !ok {
		b.saving[name] = 0
	}
	b.saving[name] += amount
}

// Withdraw 取款，返回实际取到的金额
func (b *Bank) Withdraw(name string, amount int) int {
	b.Lock()
	defer b.Unlock()
	fmt.Println("开始取钱 取", amount, "元")

	if _, ok := b.saving[name]; !ok {
		return 0
	}
	if b.saving[name] < amount {
		amount = b.saving[name]
	}
	b.saving[name] -= amount

	return amount
}

// Query 查询余额
func (b *Bank) Query(name string) int {
	b.Lock()
	defer b.Unlock()
	fmt.Println("开始")

	if _, ok := b.saving[name]; !ok {
		return 0
	}

	return b.saving[name]
}

func main() {
	b := NewBank()
	go b.Deposit("xiaoming", 100)
	go b.Withdraw("xiaoming", 20)
	go b.Deposit("xiaogang", 2000)

	time.Sleep(time.Second * 3)
	fmt.Printf("xiaoming has: %d\n", b.Query("xiaoming"))
	fmt.Printf("xiaogang has: %d\n", b.Query("xiaogang"))
}
