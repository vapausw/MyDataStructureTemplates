package main

// 栈是常用的一种线性数据结构，通常是后进入先出
// go语言一般采用数组模拟栈

type Stack struct {
	q []any
}

// 此处展示一些栈的用法
// 初始化创建一个栈
func newStack() Stack {
	return Stack{
		q: make([]any, 0, 1e9+7), //容量可以自己定义
	}
}

// Back 返回栈顶的元素
func (s *Stack) Back() any {
	if s.Empty() {
		return -1 // 栈为空值时弹出-1
	}
	return s.q[len(s.q)-1]
}

// Pop 弹出栈顶元素并返回
func (s *Stack) Pop() any {
	if s.Empty() {
		return -1 // 栈为空值时弹出-1
	}
	v := s.q[len(s.q)-1]
	s.q = s.q[:len(s.q)-1]
	return v
}

// Push 向栈的末尾加入元素
func (s *Stack) Push(v any) {
	s.q = append(s.q, v)
}

// Empty 判断当前栈是否为空
func (s *Stack) Empty() bool {
	if len(s.q) > 0 {
		return false
	}
	return true
}

// Size 返回栈中目前元素的个数
func (s *Stack) Size() int {
	return len(s.q)
}
