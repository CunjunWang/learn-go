package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(
				time.Duration(rand.Intn(1500)) *
					time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func worker(id int, c chan int) {
	// 或 for n := range c {
	for n := range c {
		time.Sleep(time.Second)
		fmt.Printf("Worker %d received %d\n", id, n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

// select 调度, 非阻塞式获取

// 程序中有c1, c2两个输入的goroutine
// worker 1个输出的goroutine
// select 1个调度的goroutine
// 四者间通过channel来通信, 不需要锁, wait之类的机制
func main() {
	// var c1, c2 chan int // c1 and c2 = nil
	var c1, c2 = generator(), generator()
	var worker = createWorker(0)

	// 从channel获取value之后放进slice里排队
	var values []int

	tm := time.After(10 * time.Second)

	tick := time.Tick(time.Second)

	for {
		var activeWorker chan<- int
		var activeValue int
		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}
		if len(values) > 10 {
			fmt.Println("Queue too long")
			return
		}
		select {
		case n := <-c1:
			values = append(values, n)
		case n := <-c2:
			values = append(values, n)
		case activeWorker <- activeValue:
			values = values[1:]
		case <-time.After(800 * time.Millisecond):
			fmt.Println("timeout")
		case <-tick:
			fmt.Println("Queue length = ", len(values))
		case <-tm:
			fmt.Println("Ends")
			return
		}
	}
}
