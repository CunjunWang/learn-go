package main

import "fmt"

// golang 指针不能运算, 不像C

// 参数传递
// pass by value / pass by reference?
// C/C++: both
// Java/python: pass by reference in most cases, except for some system inherited type
//
// C++ example:
//
// void pass_by_val(int a) {
// 	 a++;
// }
//
// void pass_by_ref(int& a) {
//	 a++;
// }
//
// int main() {
//	 int a = 3;
//
//	 pass_by_val(a);  // 3 虽然a++, 但是main里面的int没有变, 从main把a的值copy一份传入函数
//	 pass_by_ref(a);  // 4 直接把main里的a传入函数, 引用同一个变量
// }

// golang只有pass by value一种方式

// 无效, 因为是pass by value, 在func内部的操作不会反应到入参上
func swap(a, b int) {
	b, a = a, b
}

// 有效, a b 指向的位置互换
func swapPointer(a, b *int) {
	*b, *a = *a, *b
}

// "妈蛋居然没想到这样玩"
func swapReturn(a, b int) (int, int) {
	return b, a
}

func main() {
	a, b := 3, 4
	swap(a, b)
	fmt.Println(a, b)

	a2, b2 := 3, 4
	swapPointer(&a2, &b2)
	fmt.Println(a2, b2)

	a3, b3 := 3, 4
	a3, b3 = swapReturn(a3, b3)
	fmt.Println(a3, b3)
}
