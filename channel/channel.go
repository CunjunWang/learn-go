package main

import (
	"fmt"
	"time"
)

// channel: goroutine 和 goroutine之间的交互

func worker(id int, c chan int) {
	// 或 for n := range c {
	for {
		n, ok := <-c
		if !ok {
			break
		}
		fmt.Printf("Worker %d received %c\n", id, n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func channelDemo() {
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}
	// var c chan int // channel of int, c == nil
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}
	//n := <- c
	//fmt.Println(n)
	time.Sleep(time.Millisecond)
}

func bufferedChannel() {
	c := make(chan int, 3)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	time.Sleep(time.Millisecond)
}

// 关闭channel, 发送方来close, 告诉接收方没有新数据了
func channelClose() {
	c := make(chan int, 3)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c)
	time.Sleep(time.Millisecond)

}

func main() {
	// channelDemo()
	// bufferedChannel()
	channelClose()
}

// Communication Sequential Process (CSP Model)
