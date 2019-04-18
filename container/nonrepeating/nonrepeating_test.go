package main

import "testing"

func TestSubstring(t *testing.T) {
	tests := []struct {
		s   string
		ans int
	}{
		{"abcabcbb", 3},
		{"", 0},
		{"b", 1},
		{"ccc", 1},
		{"osadfuip", 8},
	}

	for _, tt := range tests {
		if actual := lengthOfNonRepeatingSubStr(tt.s); actual != tt.ans {
			t.Errorf("获取[%s]最长子串失败, 期望[%d], 实际[%d]",
				tt.s, tt.ans, actual)
		}
	}
}

// 测试性能: go test -bench .
func BenchmarkSubstring(b *testing.B) {
	s, ans := "osadfuip", 8
	for i := 0; i < 13; i++ {
		s = s + s
	}

	b.Logf("len(s) = %d", len(s))

	// 把前面准备数据的时间去掉
	b.ResetTimer()

	// 系统会自动指定次数
	for i := 0; i < b.N; i++ {
		if actual := lengthOfNonRepeatingSubStr(s); actual != ans {
			b.Errorf("获取[%s]最长子串失败, 期望[%d], 实际[%d]",
				s, ans, actual)
		}
	}
}

// 获得coverage输出: go test -coverprofile=c.out
// 获得读取c.out的UI界面: go tool cover -html=c.out

// 查看cpu使用情况 go test -bench . -cpuprofile=cpu.out
// 查看该文件: go tool pprof cpu.out
