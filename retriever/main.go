package main

import (
	"fmt"
	"learngo/retriever/mock"
	"learngo/retriever/real"
	"time"
)

// duck typing:
// 像鸭子走路，像鸭子叫，长得像鸭子，那就是鸭子
// 描述事物的外部行为而非内部结构
//
// python中的duck typing:
// def download(retriever):
//     return retriever.get("www.baidu.com")
// 运行时才知道传入的retriever有没有get
// 需要注释来说明接口
//
// c++中的duck typing:
// template <class R>
// string download(const R& retriever) {
// 		return retriever.get("www.baidu.com")
// }
// 编译时才知道传入的retriever有没有get
// 需要注释来说明接口
//
// java中没有duck typing, 有类似的代码
// <R extends Retriever>
// String download(R r) {
// 		return r.get("www.baidu.com")
// }
// 传入r, 由于R继承Retriever, 那么必须要实现get()方法
// 不是duck typing (因为必须有duck的内在属性)
//
// golang中的duck typing:
// 同时需要实现两个接口怎么办(java做不到)
// 同时具有python, c++的灵活性
// 又具有java的类型检查

// golang的接口由使用者定义

// 接口的实现是隐式的
// 只要实现接口里的方法

const url = "http://www.baidu.com"

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

func download(r Retriever) string {
	return r.Get(url)
}

func post(poster Poster) {
	poster.Post(url,
		map[string]string{
			"name": "test",
		})
}

// 接口的组合
type RetrieverPoster interface {
	Retriever
	Poster
	Connect(host string)
}

func session(session RetrieverPoster) string {
	session.Post(url, map[string]string{
		"contents": "another fake url",
	})
	return session.Get(url)
}

func main() {
	var r Retriever

	retriever := mock.Retriever{
		"this is fake baidu.com",
	}
	r = &retriever
	inspect(r)

	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	inspect(r)

	// Type assertion
	realRetriever := r.(*real.Retriever)
	fmt.Println(realRetriever.TimeOut)

	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("not a mock retriever")
	}

	fmt.Println("Try a session")
	fmt.Println(session(&retriever))
	// fmt.Println(download(r))
}

func inspect(r Retriever) {
	fmt.Printf("%T %v\n", r, r)
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents: ", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent: ", v.UserAgent)
	}
}

// 接口变量里面有什么:
// 实现者的类型 + 实现者的值
// 或者 实现者的类型 + 实现者的指针 -> 实现者
//
// 接口变量自带指针
// 采用值传递, 几乎不需要使用接口的指针
// 指针接收者必须以指针形式使用
// 值接收者都可以
