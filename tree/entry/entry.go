package main

import (
	"fmt"
	"learngo/tree"
)

// 包和封装

// 封装:
// 名字一般用CamelCase
// 首字母大写: public
// 首字母消息: private
// 所有的方法、结构、常量都遵循这一规则

// 包:
// public和private是针对包而言的
// 每个目录一个包
// 不同于java, 包名和目录名可以不一样
// main包包含了可执行入口
// 为结构定义的方法必须在同一个包内
// 但是可以是不同的文件

type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrderTraverse() {
	if myNode == nil || myNode.node == nil {
		return
	}
	// 拿到左子树 myNode.node.Left
	// 用自己定义的对象包装一下左子树
	leftNode := myTreeNode{myNode.node.Left}
	rightNode := myTreeNode{myNode.node.Right}
	leftNode.postOrderTraverse()
	rightNode.postOrderTraverse()
	myNode.node.Print()
}

func main() {

	// 创建结构体
	var root tree.Node
	fmt.Println(root)
	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	// root是结构体实例, root.right是一个指针
	// 在其他语言, 例如c++中, 要写成root.right -> left
	// 在go中也可以直接写root.right.left
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)

	root.Print()
	fmt.Println()

	//// func (node treeNode) setValue(value int) 这样定义setValue, 打印出来依然是0
	//// 因为是值传递, 改不了里面的值的
	//// 用 func (node *treeNode) setValue(value int) 传地址才行
	//root.right.left.setValue(4)
	//root.right.left.print()
	//
	//var pRoot *treeNode
	//pRoot.setValue(200)
	//pRoot = &root
	//pRoot.setValue(300)
	//pRoot.print()
	//
	//// 赋值给slice
	//nodes := []treeNode{
	//	{value: 3}, {}, {6, nil, &root},
	//}
	//fmt.Println(nodes)

	root.Traverse()
	fmt.Println()

	myRoot := myTreeNode{&root}
	myRoot.postOrderTraverse()
	fmt.Println()

	// golang没有构造函数, 可以自己写工厂函数
}
