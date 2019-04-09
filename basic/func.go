package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

// golang 函数可以返回多个值
// 通常返回 (结果, 错误)

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		// 只接收一个返回值 用下划线代替不要的返回值
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf("unsupported operation: %s" + op)
	}
}

func div(a, b int) (q, r int) {
	return a / b, a % b
}

// 函数套函数, golang是函数式变成
func apply(op func(int, int) int, a, b int) int {
	// 通过反射获取函数名
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args (%d, %d)\n", opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

// 可变参数列表
// 没有默认参数, 可选参数等功能
func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s = s + numbers[i]
	}
	return s
}

func main() {
	if result, err := eval(3, 4, "x"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}

	fmt.Println(eval(3, 4, "x"))
	// 用q r分别接收返回值
	q, r := div(13, 3)
	fmt.Println(q, r)

	// 匿名函数, 没有lambda expression
	fmt.Println(apply(
		func(a, b int) int {
			return int(math.Pow(float64(a), float64(b)))
		}, 3, 4))

	fmt.Println(sum(1, 2, 3, 4, 5))
}
