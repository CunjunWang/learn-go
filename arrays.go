package main

import "fmt"

// 定义数组
func main() {
	var arr1 [5]int                  // 默认全是o
	arr2 := [3]int{1, 3, 5}          // 用冒号定义需要给初值
	arr3 := [...]int{2, 4, 6, 8, 10} // 可以不定义长度

	var grid [4][5]int // 4行5列

	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	// 遍历数组
	// for i := 0; i < len(arr3); i++ {
	// 	 fmt.Println(arr3[i])
	// }
	// 一般用range()
	for i := range arr3 {
		fmt.Println(arr3[i])
	}
	// 可以同时获得下标和值
	for _, v := range arr3 {
		fmt.Println(v)
	}

	fmt.Println("printArray(arr1)")
	printArray(arr1[:])

	fmt.Println("printArray(arr3)")
	printArray(arr3[:])

	// 在函数内部[0]被改变了, 但是在外面, 数组本身是不变的
	fmt.Println("arr1 and arr3")
	fmt.Println(arr1, arr3)

	// 传递指针

	fmt.Println("printArray(*arr1)")
	printArrayRef(&arr1)

	fmt.Println("printArray(*arr3)")
	printArrayRef(&arr3)

	// 在函数内部[0]被改变了, 但是在外面, 数组本身是不变的
	fmt.Println("arr1 and arr3")
	fmt.Println(arr1, arr3)

}

// 数组是值类型
// [10]int 和 [20]int 是不同的
// 调用函数 func f(arr [10]int)会copy数组
// golang中一般不直接使用数组
func printArray(arr []int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func printArrayRef(arr *[5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}
