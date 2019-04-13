package tree

import "fmt"

// 定义结构体
// type 名字 struct { }
type Node struct {
	Value       int
	Left, Right *Node
}

// 为结构体定义方法, 是定义在结构体外面的
// 在函数名前面添加了接收者
// 表示这个函数print是给treeNode类型使用的
// 其实这个写法和 print(node Node)是一样的
// 只是如果用 func (node Node) print() 写法, 调用时是
// root.print()形式, 看起来更像面向对象
// 如果用 func print(node Node) 写法, 调用时就是
// print(root)了
// 依然是值传递参数. golang中所有传参都是值传递
func (node Node) Print() {
	fmt.Println(node.Value, " ")
}

// golang没有this指针, 没有self指针
// 只有使用指针才可以改变结构体内容
// nil指针也能调用方法!
func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("Setting value to nil, ignored.")
		return
	}
	node.Value = value
}

// 工厂函数, 返回的是局部变量的地址!
// c++中, 局部变量分配在栈上, 函数退出后局部变量会立刻被销毁
// java中, 几乎所有的变量都是分配在堆上的, 用new, 有gc机制
// golang不需要知道变量分配在堆还是栈上, golang编译器会自己决定的
func CreateNode(value int) *Node {
	// 坑：在c++中, 这样写程序就挂了
	// 因为返回的是一个局部变量treeNode的地址给外部调用
	// 在golang中, 这样写没有问题
	return &Node{Value: value}
}

// 值接收者 VS 指针接收者
// 要改变内容必须用指针接收者
// 结构体过大也要考虑用指针接收者
// 一致性: 如果有指针接收者, 最好都是指针接收者

// go语言特有值接收者
// 值/指针接收者都可以接收值/指针
