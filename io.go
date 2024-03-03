package main

// 此处记录常用算法模板，基本用于codeforces

import (
	"bufio"
	. "fmt"
	"math"
	"os"
	"runtime/debug"
)

func init() { debug.SetGCPercent(-1) } // 关闭垃圾收集

func solve() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	for Fscan(in, &t); t > 0; t-- {

	}
}

func memset(a []bool, v bool) []bool { //数组数据的初始化
	for i := range a {
		a[i] = v
	}
	return a
}

func unique(a []int) []int { //数组的去重
	mp := map[int]bool{}
	res := make([]int, 0, len(a))
	for _, v := range a {
		if !mp[v] {
			mp[v] = true
			res = append(res, v)
		}
	}
	return res
}
func reverse(a []int) []int { // 数组反转
	for i := 0; i < len(a)/2; i++ {
		a[i], a[len(a)-1-i] = a[len(a)-1-i], a[i]
	}
	return a
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
} // 求最小值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
} // 求最大值
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}                      // 求最大公约数
func lcm(a, b int) int { return a / gcd(a, b) * b } // 求最小公倍数
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
} // 求绝对值
// func C(n, m int) int { // 计算组合数方法一
//
//		res := make([][]int, n+1)
//		for i := range res {
//			res[i] = make([]int, m+1)
//			res[i][0] = 1
//			if i <= m {
//				res[i][i] = 1
//			}
//		}
//		for i := 1; i <= n; i++ {
//			for j := 1; j <= m; j++ {
//				res[i][j] = res[i-1][j] + res[i-1][j-1]
//			}
//		}
//		return res[n][m]
//	}

func C(n, m int) int { // 计算组合数方法二
	ans := 1
	for i := 1; i <= m; i++ {
		ans = ans * (n - m + i) / i
	}
	return ans
}

func log(a, b int) float64 { //以a为底的对数求取
	return math.Log(float64(b)) / math.Log(float64(a))
}

func main() {
	solve()
}

// 自定义类型堆
// type hp []int
//
// func (h hp) Len() int           { return len(h) }
// func (h hp) Less(i, j int) bool { return h[i] < h[j] } // > 为最大堆
// func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
// func (h *hp) Push(v any)        { *h = append(*h, v.(int)) }
// func (h *hp) Pop() any          { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
