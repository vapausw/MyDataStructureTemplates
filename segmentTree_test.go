package main

import (
	"testing"
)

// 测试普通线段树，维护的是区间最小值
func TestSegmentTree(t *testing.T) {
	a := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	st := newSegmentTree(a)

	// 测试 build
	if st[1].val != 1 {
		t.Errorf("Expected initial minimum value to be 1, got %v", st[1].val)
	}

	// 测试单点更新
	st.update(1, 5, 6) // 设置第5个元素为6
	if st.query(1, 5, 5) != 6 {
		t.Errorf("Expected value at position 5 to be 6, got %v", st.query(1, 5, 5))
	}

	// 测试区间查询
	minVal := st.query(1, 1, 3)
	if minVal != 1 {
		t.Errorf("Expected minimum value in range [1, 3] to be 1, got %v", minVal)
	}

	// 测试 queryFirstLessPos
	if pos := st.queryFirstLessPos(1, 2); pos != 1 {
		t.Errorf("Expected first position with value less than 2 to be 1, got %v", pos)
	}

	// 测试 queryFirstLessPosInRange
	if pos := st.queryFirstLessPosInRange(1, 1, 10, 2); pos != 1 {
		t.Errorf("Expected first position with value less than 2 in range [1, 10] to be 1, got %v", pos)
	}
}

// 测试懒惰线段树，维护的是区间和
func TestLazySegmentTree(t *testing.T) {
	a := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	lst := newLazySegmentTree(a)

	// 测试 build
	if lst[1].val != 110 {
		t.Errorf("Expected initial sum of lazy segment tree to be 110, got %v", lst[1].val)
	}

	// 测试区间更新和查询,此处默认查询区间为1 <= l <= r <= n，但数组a是从0开始的
	lst.update(1, 1, 5, 10)   // 区间[1, 5]每个元素增加10
	sum := lst.query(1, 1, 5) // 查询的是前5元素的和
	if sum != 80 {
		t.Errorf("Expected sum of range [1, 5] after update to be 80, got %v", sum)
	}

	// 测试 spreadAll
	lst.spreadAll(1)
	if lst[1].val != 160 {
		t.Errorf("Expected sum after spreadAll to be 160, got %v", lst[1].val)
	}
}
