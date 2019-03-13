package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// for 循环不需要括号, 可以没有初始条件
// 其他都与别的语言一样

func convertToBinary(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	// 结束条件也可以省略
	// golang里没有while, 用for可以替代while
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func forever() {
	// 死循环
	for {
		fmt.Println("abc")
	}
}

func main() {
	fmt.Println(
		convertToBinary(5),  // 101
		convertToBinary(13), // 1101
		convertToBinary(461730149),
		convertToBinary(0),
	)

	printFile("abc.txt")
}
