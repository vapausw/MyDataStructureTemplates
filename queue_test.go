package main

import (
	"testing"
)

func TestQueue(t *testing.T) {
	q := newQueue()

	// 测试新创建的队列是否为空
	if !q.Empty() {
		t.Errorf("Expected new queue to be empty")
	}

	// 测试 Push 和 Size 方法
	q.Push(1)
	q.Push(2)
	if q.Size() != 2 {
		t.Errorf("Expected size to be 2, got %v", q.Size())
	}

	// 测试 Front 方法
	if q.Front() != 1 {
		t.Errorf("Expected front to be 1, got %v", q.Front())
	}

	// 测试 Pop 方法
	if q.Pop() != 1 {
		t.Errorf("Expected Pop to return 1")
	}
	if q.Pop() != 2 {
		t.Errorf("Expected Pop to return 2")
	}

	// 测试队列是否为空
	if !q.Empty() {
		t.Errorf("Expected queue to be empty after pops")
	}

	// 测试 Pop 和 Front 方法在空队列上的行为
	if q.Pop() != -1 {
		t.Errorf("Expected Pop to return -1 for empty queue")
	}
	if q.Front() != -1 {
		t.Errorf("Expected Front to return -1 for empty queue")
	}
}
