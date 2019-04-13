package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// // 函数外也能定义变量
// var aa = 3
// var ss = "kkk"
// // 函数外定义变量不可以用:=
// // bb := true
// var bb = true
// // 这些不是全局变量，而是包内部变量

var (
	aa = 3
	ss = "kkk"
	bb = true
)

func variableZeroValue() {
	// go中, 变量名在前, 变量类型在后
	var a int
	var s string
	// %q表示quotation
	fmt.Printf("%d %q\n", a, s)
}

// go语言的变量一旦定义了就一定要用到, 不可以不用
func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

// go可以推断变量类型
func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

// := 效果和var一样
// 之后再赋值就不能用:=了
// var能不用就不用
func variableShorter() {
	a, b, c, s := 3, 4, true, "def"
	b = 5
	fmt.Println(a, b, c, s)
}

func euler() {
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))

	// 发现不是正好为0, 因为实部虚部都是浮点数, 并不准
	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)

	fmt.Println("%.3f", cmplx.Exp(1i*math.Pi)+1)
}

// 强制类型转换, 必须显示转换
func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

// 常量
// 在go语言中不会吧常量名全部大写
// 首字母大写在go中表示public
// const数值可以作为各种类型使用
func consts() {
	const filename = "abc.txt"
	const a, b = 3, 4
	const ()
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(filename, c)
}

// 枚举类型
func enums() {
	//const (
	//	cpp = 0
	//	java = 1
	//	python = 2
	//	golang = 3
	//)
	// => 简化, iota表示这一组const是自增长的
	// iota作为自增值的种子
	const (
		cpp = iota
		java
		python
		golang
	)

	// b, kb, mb, gb, tb, pb
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp, java, python, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	fmt.Println("Hello world")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, ss, bb)
	euler()
	triangle()
	consts()
	enums()
}

// 变量类型:
// bool, string
// (u)int, (u)int8, (u)int16, (u)int32, (u)int64, uintptr
// 不加u是有符号整数
// byte 8位, rune(字符型, 类似于char, 32位)
// float32, float64, complex64(实部虚部32位), complex128(实部虚部64位) 复数
