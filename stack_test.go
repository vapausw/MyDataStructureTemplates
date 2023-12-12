package main

import (
	"testing"
)

func TestStack(t *testing.T) {
	s := newStack()

	// 测试新创建的栈是否为空
	if !s.Empty() {
		t.Errorf("Expected new stack to be empty")
	}

	// 测试 Push 和 Size 方法
	s.Push(1)
	s.Push(2)
	if s.Size() != 2 {
		t.Errorf("Expected size to be 2, got %v", s.Size())
	}

	// 测试 Back 方法
	if s.Back() != 2 {
		t.Errorf("Expected back to be 2, got %v", s.Back())
	}

	// 测试 Pop 方法
	if s.Pop() != 2 {
		t.Errorf("Expected Pop to return 2")
	}
	if s.Pop() != 1 {
		t.Errorf("Expected Pop to return 1")
	}

	// 测试栈是否为空
	if !s.Empty() {
		t.Errorf("Expected stack to be empty after pops")
	}

	// 测试 Pop 和 Back 方法在空栈上的行为
	if s.Pop() != -1 {
		t.Errorf("Expected Pop to return -1 for empty stack")
	}
	if s.Back() != -1 {
		t.Errorf("Expected Back to return -1 for empty stack")
	}
}
