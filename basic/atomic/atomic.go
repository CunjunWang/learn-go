package main

import (
	"fmt"
	"sync"
	"time"
)

// mutex 互斥量

type atomicInt struct {
	value int
	lock  sync.Mutex
}

func (a *atomicInt) increment() {
	fmt.Println("Safe increment")
	func() {
		a.lock.Lock()
		// 用defer去解锁
		defer a.lock.Unlock()
		a.value++
		// 不要写在这里 a.lock.Unlock()
	}()
}

func (a *atomicInt) get() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.value
}

func main() {
	var a atomicInt
	a.increment()
	go func() {
		a.increment()
	}()
	time.Sleep(time.Millisecond)
	fmt.Println("a: ", a.get())
}
