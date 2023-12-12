package main

import (
	"testing"
)

func TestUnionFind(t *testing.T) {
	// 创建一个新的 UnionFind 实例
	uf := NewUnionFind(10)

	// 测试初始时每个节点是自己的父节点
	for i := 1; i <= 10; i++ {
		if uf.Find(i) != i {
			t.Errorf("Expected Find(%d) to be %d, got %d", i, i, uf.Find(i))
		}
	}

	// 测试合并操作
	uf.Merge(1, 2)
	if !uf.Same(1, 2) {
		t.Errorf("Nodes 1 and 2 should be in the same set after merge")
	}

	// 测试连通分量个数
	if uf.Groups != 9 {
		t.Errorf("Expected 9 groups after one merge, got %d", uf.Groups)
	}

	// 测试不在同一个集合的节点
	if uf.Same(1, 3) {
		t.Errorf("Nodes 1 and 3 should not be in the same set")
	}

	// 测试多个合并操作
	uf.Merge(2, 3)
	uf.Merge(3, 4)
	if !uf.Same(1, 4) {
		t.Errorf("Nodes 1 and 4 should be in the same set after multiple merges")
	}

	// 测试合并已经在同一集合的元素
	if uf.Merge(1, 4) != -1 {
		t.Errorf("Merge should return -1 when merging nodes already in the same set")
	}
}
