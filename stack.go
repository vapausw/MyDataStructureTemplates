package main

// 栈是常用的一种线性数据结构，通常是后进入先出
// go语言一般采用数组模拟栈
type stack struct {
	warehouse []any
}

// 此处展示一些栈的用法
// 初始化创建一个栈
func newStack() stack {
	return stack{
		warehouse: make([]any, 0, 1e9+7), //容量可以自己定义
	}
}

// Pop 弹出栈顶元素
func (s *stack) Pop() any {
	if s.Empty() {
		return -1 // 栈为空值时弹出-1
	}
	v := s.warehouse[len(s.warehouse)-1]
	s.warehouse = s.warehouse[:len(s.warehouse)-1]
	return v
}

// Push 向栈的末尾加入元素
func (s *stack) Push(v any) {
	s.warehouse = append(s.warehouse, v)
}

// Empty 判断当前栈是否为空
func (s *stack) Empty() bool {
	if len(s.warehouse) > 0 {
		return false
	}
	return true
}

// Size 返回栈中目前元素的个数
func (s *stack) Size() int {
	return len(s.warehouse)
}
