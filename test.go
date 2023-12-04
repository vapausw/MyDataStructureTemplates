package main

func main() {
	roads := [][]int{{0, 1}, {0, 2}, {1, 3}, {1, 4}}
	minimumFuelCost(roads, 5)
}

func minimumFuelCost(roads [][]int, seats int) int64 {
	n := len(roads)
	g := make([][]int, n+1)
	for _, e := range roads {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	ans := 0
	var dfs func(int, int) (int, int) //返回车辆数以及目前所有的乘客数量
	dfs = func(x, fa int) (int, int) {
		flag := true
		res, cnt := 0, 0
		for _, y := range g[x] {
			if y != fa {
				flag = false
				res1, cnt1 := dfs(y, x)
				res += res1
				cnt += cnt1
			}
		}
		ans += res
		if flag {
			return 1, 1
		}
		res = (cnt + seats) / seats // 当前乘客数量合并,向上取整
		return res, cnt + 1
	}
	dfs(0, -1)
	return int64(ans)
}
