package main

import (
	"bufio"
	. "fmt"
	"os"
)

// 此文件用于存储一些题目的小技巧
// 此处代码基本使用类似与LeetCode https://leetcode.cn/ 的代码风格，不使用完整输入输出，只使用返回值完成题目

// 题目链接 https://leetcode.cn/problems/minimum-fuel-cost-to-report-to-the-capital/?envType=daily-question&envId=2023-12-05
// 此题目所使用的技巧是，图的建立以及图的遍历和贪心的思想
func minimumFuelCost(roads [][]int, seats int) int64 {
	n := len(roads)
	g := make([][]int, n+1)
	for _, e := range roads { //图的建立，无向图
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	ans := 0
	var dfs func(int, int) (int, int)  //返回车辆数以及目前所有的乘客数量， 贪心的思想
	dfs = func(x, fa int) (int, int) { // 图的遍历，深度优先搜索
		res, cnt := 0, 0
		for _, y := range g[x] {
			if y != fa {
				res1, cnt1 := dfs(y, x)
				res += res1
				cnt += cnt1
			}
		}
		if res == 0 {
			return 1, 1
		}
		ans += res
		res = (cnt + seats) / seats // 当前乘客数量合并,使所使用的车辆最小,向上取整
		return res, cnt + 1
	}
	dfs(0, -1)
	return int64(ans)
}

// 题目链接 https://codeforces.com/contest/1907/problem/C
// 一道比较新的二分check思路，需要检查区间可达位置，每次更新当前拥有的区间范围就可

func Cf_913_D() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush() // 加速读取
	var t, n int
	type pair struct {
		li, ri int
	}
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n)
		a := make([]pair, n)
		for i := range a {
			Fscan(in, &a[i].li, &a[i].ri)
		}
		var check func(int) bool
		check = func(x int) bool {
			left, right := 0, 0
			for i := 0; i < n; i++ {
				if right+x < a[i].li || left-x > a[i].ri {
					return false
				}
				left = max(a[i].li, left-x)
				right = min(a[i].ri, right+x)
			}
			return true
		}
		l, r := -1, int(1e9+1)
		for l+1 < r {
			mid := l + (r-l)>>1
			if check(mid) {
				r = mid
			} else {
				l = mid
			}
		}
		Fprintln(out, r)
	}
}

// 一个关于删除相邻且不相同元素剩余个数的结论题目，寻找当前字符串中的元素相同最多的元素个数记为maxv，如果maxv小于等于n / 2 则所有元素都可以抵消，根据n是奇数还是偶数
// 来判定最后剩余一个还是两个，如果maxv > n / 2， 则最后剩余的元素个数就是maxv所不能自己消除的元素maxv - (n - maxv)
// 题目链接 https://codeforces.com/contest/1907/problem/C

func Cf_913_C() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush() // 加速读取
	var t, n int
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n)
		var s []byte
		Fscan(in, &s)
		res := map[byte]int{}
		for _, c := range s {
			res[c]++
		}
		ans, maxv := 0, 0
		for _, v := range res {
			maxv = max(maxv, v)
		}
		if n%2 != 0 {
			ans = 1
		}
		if maxv >= (n+1)/2 {
			ans = maxv - (n - maxv)
		}
		Fprintln(out, ans)
	}
}
