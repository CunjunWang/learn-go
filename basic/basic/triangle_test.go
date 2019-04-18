package main

import "testing"

// 表格驱动测试
// 当前目录测试: go test .

func TestTriangle(t *testing.T) {
	tests := []struct{ a, b, c int }{
		{3, 4, 5},
		{5, 12, 13},
		{8, 15, 17},
		{12, 35, 37},
		{30000, 40000, 50000},
	}

	for _, tt := range tests {
		if actual := calculateTriangle(tt.a, tt.b); actual != tt.c {
			t.Errorf("计算三角形(%d, %d); "+"获得 %d, 期望 %d",
				tt.a, tt.b, actual, tt.c)
		}
	}
}
