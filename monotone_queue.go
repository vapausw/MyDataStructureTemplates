package main

// 以下是单调队列模板可以直接复制使用

type MonotoneQueue struct {
	Data []any
	Size int //单调队列大小
}

// Less 区间维护逻辑函数
func (mq *MonotoneQueue) Less(a, b any) bool {
	return a.(int) > b.(int) // >= 维护区间最大值，单调递增队列， <= 维护区间最小值，单调递减队列， =判断队列中是否允许重复值存在
}

// PushPurge  清除队列中不合法的元素，写此函数是因为有时需要更新之前计算答案
func (mq *MonotoneQueue) PushPurge(v any) {
	for mq.Size > 0 && mq.Less(v, mq.Data[mq.Size-1]) {
		mq.Data = mq.Data[:mq.Size-1]
		mq.Size--
	}
}

// Push 向单调队列中添加元素
func (mq *MonotoneQueue) Push(v any) {
	mq.Size++
	mq.Data = append(mq.Data, v)
}

// Pop 删除单调队列的第一个元素并返回
func (mq *MonotoneQueue) Pop() any {
	if mq.Size <= 0 {
		return -1 //非法操作
	}
	v := mq.Data[0]
	mq.Data = mq.Data[1:]
	mq.Size--
	return v
}

// Top 返回单调队列的第一个元素
func (mq *MonotoneQueue) Top() any {
	if mq.Size <= 0 {
		return -1
	}
	return mq.Data[0]
}

// 单调队列模板题目 https://leetcode.cn/problems/dui-lie-de-zui-da-zhi-lcof/

type Checkout struct {
	values []int
	q      MonotoneQueue
}

func Constructor() Checkout {
	return Checkout{}
}

func (this *Checkout) Get_max() int {
	if this.q.Size == 0 {
		return -1
	}
	return this.q.Top().(int)
}

func (this *Checkout) Add(value int) {
	this.values = append(this.values, value)
	this.q.PushPurge(value)
	this.q.Push(value)
}

func (this *Checkout) Remove() int {
	if len(this.values) == 0 {
		return -1
	}
	temp := this.values[0]
	this.values = this.values[1:]
	if temp == this.q.Top().(int) {
		return this.q.Pop().(int)
	}
	return temp
}

// 单调队列题目 https://leetcode.cn/problems/shortest-subarray-with-sum-at-least-k/description/

// 此处先使用前缀和方便计算子数组的值
func shortestSubarray(nums []int, k int) int {
	n := len(nums)
	s := make([]int, n+1)
	for i, x := range nums {
		s[i+1] = s[i] + x
	}
	ans := n + 1
	var q []int
	for i, curS := range s {
		for len(q) > 0 && curS-s[q[0]] >= k { // 子数组的值大于等于k，更新答案
			ans = min(ans, i-q[0])
			q = q[1:]
		}
		for len(q) > 0 && s[q[len(q)-1]] >= curS { // 此处保证单调队列是单调递增的，且允许有重复元素
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}
	if ans > n {
		return -1
	}
	return ans
}

// 单调栈以及单调队列都遵循一个单调的准则要不单调递增要不就单调递减，且栈是后进先出，队列是先进先出，根据具体需求选取，且根据具体需求看是否允许存在重复元素
