package main

import "fmt"

// 函数式编程 vs 函数指针
// 参数、变量、返回值都可以是函数
// 高阶函数
// 函数 -> 闭包

// "正统"的函数式编程
// 不可变: 不能有状态, 只有常量和函数
// 函数只能有一个参数

// 闭包[ 函数体[局部变量 自由变量] ]

// adder()函数没有入参, 出参是一个函数func(int) int
// 出参函数的入参是一个int, 出参是一个int
func adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, adder2(base + v)
	}
}

func main() {
	a := adder()
	// a 是adder()的返回值, 是一个带参的函数
	for i := 0; i < 10; i++ {
		// 这里就可以把i传进a
		fmt.Println(a(i))
	}
}

// python中的闭包
// def adder():
//	  sum = 0
//
//	  def f(value):
//		  nonlocal sum
//		  sum += value
//		  return sum
//	  return f
// python原生支持闭包
// 使用__closure__查看闭包内容

// c++中的闭包
// auto adder() {
// 		auto sum = 0;
// 		return [=] (int value) mutable {
// 			sum += value;
// 			return sum;
// 		}
// }
// 过去: stl或者boost带有类似的库
// c++11之后支持闭包, 这段代码在c++14下编译通过

// java中的闭包
// Function<Integer, Integer> adder() {
//	  final Holder<Integer> sum = new Holder<>(0);
//	  return (Integer value) -> {
//		  sum.value += value;
//		  return sum.value;
//	  }
// }
// 1.8以后：使用Function接口和Lambda表达式来创建函数对象
// 匿名类或Lambda表达式均支持闭包
