package main

import "fmt"

// arr := [...]int{0,1,2,3,4,5,6,7}
// s := arr[2:6] // 区间默认为半开半闭区间, 2包含, 6不包含
// s值为[2 3 4 5]

// slice is in the form of arr[]
// slice is a view of an array

func updateSlice(s []int) {
	s[0] = 100
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	fmt.Println("arr[2:6] = ", arr[2:6])
	fmt.Println("arr[:6] = ", arr[:6])
	s1 := arr[2:]
	fmt.Println("s1 = ", s1)
	s2 := arr[:]
	fmt.Println("s2 = ", s2)

	updateSlice(s1)
	fmt.Println("s1 after update = ", s1)
	fmt.Println("arr after update = ", arr)

	updateSlice(s2)
	fmt.Println("s2 after update = ", s2)
	fmt.Println("arr after update = ", arr)

	fmt.Println("Re-slice:")
	fmt.Println("s2 before re-slice = ", s2)
	s2 = s2[:5]
	fmt.Println("s2 after slice s2[:5] = ", s2)
	s2 = s2[2:]
	fmt.Println("s2 after slice s2[2:] = ", s2)

	// slice的扩展
	fmt.Println("Extending slice")
	arr[0], arr[2] = 0, 2
	s1 = arr[2:6]
	//fmt.Println(s1[4])
	s2 = s1[3:5] // s1[3], s1[4]
	fmt.Println("s1 = ", s1)
	fmt.Println("s2 = ", s2)

	// slice实现:
	// ptr, len, cap
	// ptr指向slice的头, len表示slice长度, 方括号能取到的值
	// cap表示slice对应的整个array的长度, 因此不超过cap都可以扩展
	// slice可以向后扩展, 不可以向前扩展
	// s[i]不可以超过len(s), 向后扩展不可以超越底层数组cap(s)
	// 获取len和cap:
	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d\n", s2, len(s2), cap(s2))
	fmt.Println(s1[3:6])
	// fmt.Println(s1[3:7]) // 报错, slice bound看的是cap

	// 向slice中添加元素
	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	fmt.Println("s3, s4, s5 = ", s3, s4, s5)
	// s4, s5 view a different array
	fmt.Println("arr = ", arr)
	// 添加元素时如果超过cap, 系统会重新分配更大的底层数组
	// 由于值传递, 必须接受append的返回值 s = append(s, val)
}
