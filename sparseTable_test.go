package main

import (
	"testing"
)

func TestST(t *testing.T) {
	// 定义一个示例数组
	a := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}

	// 创建 ST 表
	st := NewSt(a)

	// 定义一些测试的区间，以及预期的最大值
	tests := []struct {
		l, r, expected int
	}{
		{0, 3, 5},   // 区间 [0, 3) 的最大值
		{4, 7, 13},  // 区间 [4, 7) 的最大值
		{2, 5, 9},   // 区间 [2, 5) 的最大值
		{6, 10, 19}, // 区间 [6, 10) 的最大值
	}

	// 遍历测试
	for _, tt := range tests {
		res := st.Query(tt.l, tt.r)
		if res != tt.expected {
			t.Errorf("Query(%d, %d) = %d, want %d", tt.l, tt.r, res, tt.expected)
		}
	}
}
