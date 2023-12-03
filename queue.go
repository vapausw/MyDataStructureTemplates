package main

// 栈是常用的一种线性数据结构，通常是后进入先出
// go语言一般采用数组模拟栈
type queue struct {
	q []any
}

// 此处展示一些栈的用法
// 初始化创建一个栈
func newQueue() queue {
	return queue{
		q: make([]any, 0, 1e9+7), //容量可以自己定义
	}
}

// Pop 弹出队首元素
func (f *queue) Pop() any {
	if f.Empty() {
		return -1 // 队列为空值返回-1
	}
	v := f.q[0]
	f.q = f.q[1:]
	return v
}

// Push 向队列的末尾加入元素
func (f *queue) Push(v any) {
	f.q = append(f.q, v)
}

// Empty 判断当前队列是否为空
func (f *queue) Empty() bool {
	if len(f.q) > 0 {
		return false
	}
	return true
}

// Size 返回队列中目前元素的个数
func (f *queue) Size() int {
	return len(f.q)
}
