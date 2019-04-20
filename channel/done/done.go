package main

import (
	"fmt"
	"sync"
)

// channel: goroutine 和 goroutine之间的交互

// 不要通过共享内存来通信, 通过通信来共享内存

type worker struct {
	in   chan int
	done func()
}

func doWork(id int, w worker) {
	// 或 for n := range c {
	for n := range w.in {
		fmt.Printf("Worker %d received %c\n", id, n)
		go func() {
			w.done()
		}()
	}
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doWork(id, w)
	return w
}

func channelDemo() {
	var workers [10]worker
	var wg sync.WaitGroup
	wg.Add(20)

	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}

	// var c chan int // channel of int, c == nil
	for i, worker := range workers {
		worker.in <- 'a' + i
	}
	for i, worker := range workers {
		worker.in <- 'A' + i
	}

	// wait for all of them
	// for _, worker := range workers {
	//	  <- worker.done
	//	  <- worker.done
	// }
	wg.Wait()
}

func main() {
	channelDemo()
}

// Communication Sequential Process (CSP Model)
