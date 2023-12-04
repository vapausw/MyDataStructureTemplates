package main

// 此处代码思路参考于灵神的代码模板
// 用两个 slice 头对头拼在一起实现

// Deque 定义对列的基本结构
type Deque struct {
	l, r []any // l的尾部为双端队列的头， r的尾部为双端队列的尾部
}

// PushFront 向双端队列的队首添加元素
func (q *Deque) PushFront(v any) {
	q.l = append(q.l, v)
}

// PushBack 向双端队列的队尾添加元素
func (q *Deque) PushBack(v any) {
	q.r = append(q.r, v)
}

// Front 返回队首元素
func (q *Deque) Front() any {
	if len(q.l) > 0 {
		return q.l[len(q.l)-1]
	} else if len(q.r) > 0 {
		return q.r[0]
	}
	return -1 // 代表当前队列无元素
}

// Back 返回队尾元素
func (q *Deque) Back() any {
	if len(q.r) > 0 {
		return q.r[len(q.r)-1]
	} else if len(q.l) > 0 {
		return q.l[0]
	}
	return -1
}

// PopFront 返回双端队列的队首并且出队
func (q *Deque) PopFront() (v any) {
	if len(q.l) > 0 {
		q.l, v = q.l[:len(q.l)-1], q.l[len(q.l)-1]
	} else if len(q.r) > 0 {
		v, q.r = q.r[0], q.r[1:]
	}
	return
}

// PopBack 返回双端队列的队尾并且出队
func (q *Deque) PopBack() (v any) {
	if len(q.r) > 0 {
		q.r, v = q.r[:len(q.r)-1], q.r[len(q.r)-1]
	} else if len(q.l) > 0 {
		v, q.l = q.l[0], q.l[1:]
	}
	return
}

// Empty 返回双端队列是否为空
func (q *Deque) Empty() bool {
	return len(q.l) == 0 && len(q.r) == 0
}

// Size 返回双端队列的元素个数
func (q *Deque) Size() int {
	return len(q.l) + len(q.r)
}

// Get 0 <= i < q.Size(), 获取双端队列中的第i个元素
func (q *Deque) Get(i int) any {
	if q.Size() < i { // 不合法返回-1
		return -1
	}
	if i < len(q.l) {
		return q.l[len(q.l)-1-i]
	}
	return q.r[i-len(q.l)]
}
