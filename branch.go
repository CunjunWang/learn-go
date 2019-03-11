package main

import (
	"io/ioutil"
	"fmt"
)

// ======================= switch =========================
// switch 会自动break, 除非使用fallthrough
// panic(): 报错信息, 中断程序执行
// switch后可以没有表达式, 直接在case里写条件

func grade(score int) string {
	grade := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score: %d", score))
	case score < 60:
		grade = "F"
	case score < 80:
		grade = "C"
	case score < 90:
		grade = "B"
	case score <= 100:
		grade = "A"
	}
	return grade
}

func main() {

	// ======================= if =========================

	// if 条件里可以赋值
	// if 条件里赋值的变量作用域就在这个语句里

	const filename = "abc.txt"
	// contents, err := ioutil.ReadFile(filename)
	// if err != nil {
	//	 fmt.Println(err)
	// } else {
	//	 fmt.Printf("%s\n", contents)
	// }
	//
	// => 简化
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
	// fmt.Println(contents)

	fmt.Println(
		grade(0),
		grade(59),
		grade(60),
		grade(99),
		grade(100),
		//grade(101),
	)

}
