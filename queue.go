package main

// 队列是常用的一种线性数据结构，通常是先进先出
// go语言一般采用数组模拟队列

type Queue struct {
	q []any
}

// 此处展示一些队列的用法
// 初始化创建一个队列
func newQueue() Queue {
	return Queue{
		q: make([]any, 0, 1e9+7), //容量可以自己定义
	}
}

func (f *Queue) Front() any {
	if f.Empty() {
		return -1 // 队列为空值返回-1
	}
	return f.q[0]
}

// Pop 弹出队首元素并返回队首元素
func (f *Queue) Pop() any {
	if f.Empty() {
		return -1 // 队列为空值返回-1
	}
	v := f.q[0]
	f.q = f.q[1:]
	return v
}

// Push 向队列的末尾加入元素
func (f *Queue) Push(v any) {
	f.q = append(f.q, v)
}

// Empty 判断当前队列是否为空
func (f *Queue) Empty() bool {
	if len(f.q) > 0 {
		return false
	}
	return true
}

// Size 返回队列中目前元素的个数
func (f *Queue) Size() int {
	return len(f.q)
}
