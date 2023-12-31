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

// 题目链接 https://leetcode.cn/problems/stamping-the-grid/?envType=daily-question&envId=2023-12-14
// 此题目所使用的技巧为二维前缀和与二维差分数组
// 此解题方法参考 https://leetcode.cn/problems/stamping-the-grid/solutions/1199642/wu-nao-zuo-fa-er-wei-qian-zhui-he-er-wei-zwiu/?envType=daily-question&envId=2023-12-14
func possibleToStamp(grid [][]int, stampHeight int, stampWidth int) bool {
	m, n := len(grid), len(grid[0])
	pre := make([][]int, m+1)
	pre[0] = make([]int, n+1)
	// 求二维前缀和
	for i, row := range grid {
		pre[i+1] = make([]int, n+1)
		for j, v := range row {
			pre[i+1][j+1] = pre[i][j+1] + pre[i+1][j] - pre[i][j] + v
		}
	}
	// 二维差分数组,此处多两列为了后面还原计算方便
	diff := make([][]int, m+2)
	for i := range diff {
		diff[i] = make([]int, n+2)
	}
	for i2 := stampHeight; i2 <= m; i2++ {
		for j2 := stampWidth; j2 <= n; j2++ {
			j1 := j2 - stampWidth + 1
			i1 := i2 - stampHeight + 1
			if pre[i2][j2]-pre[i2][j1-1]-pre[i1-1][j2]+pre[i1-1][j1-1] == 0 { // 当前矩形可以放置邮票
				diff[i1][j1]++
				diff[i1][j2+1]--
				diff[i2+1][j1]--
				diff[i2+1][j2+1]++
			}
		}
	}
	// 二维前缀和进行还原差分数组
	for i, row := range grid {
		for j, v := range row {
			diff[i+1][j+1] += diff[i+1][j] + diff[i][j+1] - diff[i][j]
			if v == 0 && diff[i+1][j+1] == 0 {
				return false
			}
		}
	}
	return true
}

// 题目链接 https://codeforces.com/contest/1914/problem/F
/*
题目描述：
BerSoft 是伯兰最大的 IT 公司。BerSoft 公司有 n 名员工，编号从 1 到 n 。

第一名员工是公司负责人，他没有任何上级。其他每个员工 i都有一个直接上级 pi;

如果以下条件之一成立，则认为员工 x是员工y 的上级(直接或间接)：

- 员工 x 是员工 y的直接上级；
- 员工 x 是员工 y的直接上级的上级。

BerSoft 的组织结构是，公司领导是每位员工的上级。

即将举行编程比赛。为此应成立两人小组。但是，如果团队中的一名员工是另一名员工的上级，他们在一起就会很不自在。因此，应建立两人小组，这样就不会出现谁比谁高一等的情况。注意，任何员工都不能参加一个以上的团队。

你的任务是根据上述规则计算出团队的最大可能数量。
*/
// 学习灵神CF 916 F 的思路
// 对其分为两种情况，（1）当前根节点下的子树节点最大值不超过子树节点值和的一半， 这种可以直接算答案
// （2） 当前根节点下的子树节点最大值超过子树节点值和的一半， 取其与子树的节点与最大值的节点合并为一组，然后这个最大值节点上的子树就全部挂到了根，这样会慢慢将其转换为第一种情况
// 如此该问题就得解了

func Cf_916_F() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var t, n, v int
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n)
		g := make([][]int, n)
		for i := 1; i < n; i++ {
			Fscan(in, &v)
			v-- // 将节点值保证在0——n之间
			g[v] = append(g[v], i)
		}
		size := make([]int, n) // 记录子节点的数量
		var initSize func(int)
		initSize = func(x int) {
			size[x] = 1
			for i, y := range g[x] {
				initSize(y)
				size[x] += size[y]
				if size[y] > size[g[x][0]] { // 保证最左边是节点树最多的子树，方便后续处理
					g[x][0], g[x][i] = g[x][i], g[x][0]
				}
			}
		}
		initSize(0)

		ans, other, x := 0, 0, 0
		for {
			if other > 0 { // 此处的目的是将x的根与other进行组合也算为一种方案
				ans++
				other--
			}
			if len(g[x]) == 0 {
				break
			}

			s := size[x] - 1
			y := g[x][0]

			if size[y]*2 <= s+other { // 满足第一种情况
				ans += (s + other) / 2
				break
			}
			// 不满足，开始转化
			other += s - size[y]
			x = y
		}
		Fprintln(out, ans)
	}
}
