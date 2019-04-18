package queue

// interface{} 任何类型

type Queue []interface{}

// 改变了Queue, 要用指针接收
func (q *Queue) Push(v interface{}) {
	*q = append(*q, v)
}

func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head.(int)
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

// 写文档
// go doc xxx
// godoc -http :6060
