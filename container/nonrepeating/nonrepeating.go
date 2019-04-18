package main

import "fmt"

// abcabcbb -> abc 3
// bbbbbbbb -> b   1

var lastOccurred = make([]int, 0xffff)

func lengthOfNonRepeatingSubStr(s string) int {
	//lastOccurred := make(map[byte]int)

	// 优化, 用空间换时间, c++常用手段
	for i := range lastOccurred {
		lastOccurred[i] = -1
	}

	start := 0
	maxLength := 0

	for i, ch := range []byte(s) {
		if lastI := lastOccurred[ch]; lastI != -1 && lastI >= start {
			start = lastOccurred[ch] + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}

	return maxLength
}

func main() {
	fmt.Println(lengthOfNonRepeatingSubStr("abcabcbb"))
	fmt.Println(lengthOfNonRepeatingSubStr(""))
	fmt.Println(lengthOfNonRepeatingSubStr("b"))
	fmt.Println(lengthOfNonRepeatingSubStr("ccc"))
	fmt.Println(lengthOfNonRepeatingSubStr("osadfuip"))
}
