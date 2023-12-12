package main

import (
	"testing"
)

func TestDeque(t *testing.T) {
	q := &Deque{}

	// 测试空队列
	if !q.Empty() {
		t.Errorf("Expected deque to be empty")
	}

	// 测试 PushFront 和 Front
	q.PushFront(1)
	if q.Front() != 1 {
		t.Errorf("Expected front to be 1, got %v", q.Front())
	}

	// 测试 PushBack 和 Back
	q.PushBack(2)
	if q.Back() != 2 {
		t.Errorf("Expected back to be 2, got %v", q.Back())
	}

	// 测试 Size
	if q.Size() != 2 {
		t.Errorf("Expected size to be 2, got %v", q.Size())
	}

	// 测试 PopFront
	if q.PopFront() != 1 {
		t.Errorf("Expected PopFront to return 1")
	}

	// 测试 PopBack
	if q.PopBack() != 2 {
		t.Errorf("Expected PopBack to return 2")
	}

	// 测试空队列属性
	if !q.Empty() {
		t.Errorf("Expected deque to be empty after pops")
	}

	// 测试 Get 方法
	q.PushFront(3)
	q.PushFront(4)
	if q.Get(0) != 4 {
		t.Errorf("Expected Get(0) to return 4, got %v", q.Get(0))
	}
	if q.Get(1) != 3 {
		t.Errorf("Expected Get(1) to return 3, got %v", q.Get(1))
	}
}
