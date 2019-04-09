package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "yes测试中文!" // utf-8 编码
	fmt.Println(s)
	fmt.Println(len(s))
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Println()

	for i, ch := range s { // ch is a rune
		fmt.Printf("(%d %X) ", i, ch)
	}
	fmt.Println()

	fmt.Println("Run count = ", utf8.RuneCountInString(s))

	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c) ", i, ch)
	}
	fmt.Println()

	// rune相当于go的char
	// 使用range遍历pos, rune对
	// 使用utf8.RuneCountInString获得字符数量
	// len获得字节长度
	// []byte获得字节
}
