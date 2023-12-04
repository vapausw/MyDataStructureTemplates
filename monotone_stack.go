package main

// 单调栈模板题 https://leetcode.cn/problems/daily-temperatures/
// 解题代码如下
func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	ans := make([]int, n)
	var st []int //此处使用单调栈
	for i, t := range temperatures {
		for len(st) > 0 && t > temperatures[st[len(st)-1]] { // 栈不为空，且单调栈尾小于t，单调栈维护递增值，且元素不一样
			j := st[len(st)-1]
			st = st[:len(st)-1]
			ans[j] = i - j // 此处更新答案，更高的气温在i - j天后出现
		}
		st = append(st, i)
	}
	return ans
}

// 单调栈的思路更多的还是与贡献法或者数据结构优化题目相结合。
