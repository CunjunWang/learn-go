package main

import (
	"fmt"
	"runtime"
	"time"
)

// 并发编程
// go 关键字

func main() {
	// 死循环, 一直打印i=0的状态
	// for i := 0; i < 10; i++ {
	// 	 func(i int) {
	// 		 for {
	// 			 fmt.Printf("Hello from goroutine %d\n", i)
	// 		 }
	//   }(i)
	// }

	// 1000 并发
	// for i := 0; i < 1000; i++ {
	//    go func(i int) {
	//		  for {
	//			  fmt.Printf("Hello from goroutine %d\n", i)
	//		  }
	//	  }(i)
	//  }
	//  time.Sleep(time.Millisecond)

	var a [10]int
	for i := 0; i < 10; i++ {
		go func(i int) { // race condition
			for {
				a[i]++
				// 主动交出控制权, 让别人执行
				// 否则会卡在这里
				// io操作会主动交出控制权
				runtime.Gosched()
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
	fmt.Println(a)

	// go run -race xxx.go 查看race condition
}

// 协程 Coroutine
// 轻量级线程
// 非抢占式多任务处理, 由协程主动交出控制权
// 编译器、解释器、虚拟机层面的多任务
// 多个协程可以再一个或者多个线程上运行

// 普通函数
// 线程{ main -> doWork }
// 协程 { 可能一个线程{ main <-> doWork } }

// 其他语言中的协程

// c++: Boost.Coroutine

// Java: 不支持

// python:
// 使用yield关键字实现协程
// python3.5加入了async def原生支持协程

// goroutine

//
//
// 【 线程 		【 线程		  【 线程
// 								 goroutine
//   goroutine     goroutine     goroutine
//   goroutine				     goroutine
// 								 goroutine
// 	】			 】			   】
//
// 【                调度器                】
//
// 调度器决定几个goroutine放在一个线程执行

// 任何函数只要前面加上go就能送给调度器运行
// 不需要在定义时区分是否是异步函数
// 调度器在合适的点进行切换
// 用 -race 来查看数据冲入

// goroutine可能切换的点:
// I/O, select
// channel
// 等待锁
// 函数调用(有时)
// runtime.Gosched()
// 只是参考, 不能保证切换, 不能保证在其他地方不切换
