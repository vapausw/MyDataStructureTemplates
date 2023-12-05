package main

// 以下是单调栈模板可以直接复制使用

type MonotoneStack struct {
	Data []any
	Size int //单调栈大小
}

// Less 区间维护逻辑函数
func (mq *MonotoneStack) Less(a, b any) bool {
	return a.(int) < b.(int) // >= 维护区间最大值，单调递增队列， <= 维护区间最小值，单调递减队列， =判断队列中是否允许重复值存在
}

// PushPurge 清除栈中不合法的元素，写此函数是因为有时需要更新之前计算答案，具体看下述用例
func (mq *MonotoneStack) PushPurge(v any) {
	for mq.Size > 0 && mq.Less(v, mq.Data[mq.Size-1]) {
		mq.Pop()
	}
}

func (mq *MonotoneStack) Push(v any) {
	mq.Size++
	mq.Data = append(mq.Data, v)
}

// Pop 删除单调栈的最后一个元素并返回
func (mq *MonotoneStack) Pop() any {
	if mq.Size <= 0 {
		return -1 //非法操作
	}
	v := mq.Data[mq.Size-1]
	mq.Data = mq.Data[:mq.Size-1]
	mq.Size--
	return v
}

// Top 返回单调栈的最后一个元素
func (mq *MonotoneStack) Top() any {
	if mq.Size <= 0 {
		return -1
	}
	return mq.Data[mq.Size-1]
}

// 使用上面单调栈模板的示例 https://leetcode.cn/problems/final-prices-with-a-special-discount-in-a-shop/
func finalPrices(prices []int) []int {
	var q MonotoneStack
	ans := make([]int, len(prices))
	for i := len(prices) - 1; i >= 0; i-- {
		q.PushPurge(prices[i])
		if q.Size == 0 {
			ans[i] = prices[i]
		} else {
			ans[i] = prices[i] - q.Top().(int)
		}
		q.Push(prices[i])
	}
	return ans
}

// 手写单调栈模板题 https://leetcode.cn/problems/daily-temperatures/
// 解题代码如下，因为要在维护的过程中更新答案，而且维护的是下标不是单一的值
func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	ans := make([]int, n)
	var st []int
	for i, t := range temperatures {
		for len(st) > 0 && t > temperatures[st[len(st)-1]] {
			j := st[len(st)-1]
			st = st[:len(st)-1]
			ans[j] = i - j
		}
		st = append(st, i)
	}
	return ans
}

// 单调栈的思路更多的还是与贡献法或者数据结构优化题目相结合。
