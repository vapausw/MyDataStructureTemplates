package main

import (
	"testing"
)

func TestFenwick(t *testing.T) {
	f := NewFenwick(10)

	// 测试单点修改和查询
	f.Add(1, 5)
	if f.Query(1, 1) != 5 {
		t.Errorf("Expected Query(1, 1) to return 5, got %v", f.Query(1, 1))
	}

	f.Add(2, 3)
	if f.Query(1, 2) != 8 {
		t.Errorf("Expected Query(1, 2) to return 8, got %v", f.Query(1, 2))
	}

	// 测试二维树状数组
	fd := newFenwickDiff(10)

	// 测试区间修改和查询
	fd.add(1, 5, 2) // 将区间 [1, 5] 的每个元素加 2
	if fd.Query(1, 5) != 10 {
		t.Errorf("Expected Query(1, 5) after adding 2 to each element to return 10, got %v", fd.Query(1, 5))
	}

	fd.add(3, 4, 3) // 将区间 [3, 4] 的每个元素加 3
	if fd.Query(1, 5) != 16 {
		t.Errorf("Expected Query(1, 5) after additional updates to return 16, got %v", fd.Query(1, 5))
	}

	// 测试单点查询
	if fd.Query(3, 3) != 5 {
		t.Errorf("Expected Query(3, 3) to return 5, got %v", fd.Query(3, 3))
	}
}
