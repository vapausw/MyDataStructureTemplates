package main

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
