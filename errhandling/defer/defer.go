package main

import (
	"bufio"
	"errors"
	"fmt"
	"learngo/functional/fib"
	"os"
)

// defer:
// 不马上执行, 在函数退出前执行
// defer相当于有一个stack, 先进后出
// 参数在defer后计算

func tryDefer() {
	//defer fmt.Println(1)
	//defer fmt.Println(2)
	//fmt.Println(3)
	//panic("error occurred")
	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
		if i == 30 {
			panic("printed too many")
		}
	}
}

// 0666 unix权限, 可读可写
func writeFile(fileName string) {
	file, err := os.OpenFile(
		fileName, os.O_EXCL|os.O_CREATE, 0666)

	// 自建error
	err = errors.New("this is a custom error")

	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			// pathError.Op 操作;
			// pathError.Path 什么路径什么文件;
			// pathError.Err 什么错;
			fmt.Println(pathError.Op,
				pathError.Path,
				pathError.Err)
		}
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fibonacci()
	for i := 0; i < 50; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	writeFile("fib.txt")
	// tryDefer()
}

// 何时使用defer
// open / close
// lock / unlock
// printHeader / printFooter
