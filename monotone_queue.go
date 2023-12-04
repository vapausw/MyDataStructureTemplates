package main

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
