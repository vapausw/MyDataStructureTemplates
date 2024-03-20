package main

import (
	"bufio"
	"fmt"
	"os"
)

// 主要用于解决树上问题
// 树是没有环的连通图
// 树的特点：1. 有且仅有一个根节点 2. 除根节点外，每个节点有且仅有一个父节点 3. 有且仅有一个根节点 4. 从根节点到任意节点有且仅有一条路径

type child struct {
	index, value int // 存储当前节点的值，以及与其相连的边权值
}
type Tree struct {
	children [][]child // 存储树一般存储相连的便即可，例如1相连的边为2， 3， 4，那么g[1] = {2, 3, 4}, 使用数组模拟邻接表
}

func (t *Tree) New(n int) { // 创建一个有n个节点的树
	t.children = make([][]child, n)
}

func (t *Tree) Add(x, y, value int) { // 添加树的边，从x指向y的边， 无向图的添加方式
	t.children[x] = append(t.children[x], child{y, value})
	t.children[y] = append(t.children[y], child{x, value})
}

// DoubleLCA 倍增法求LCA，支持在线查询, 核心思想倍增法 + 二进制拆分
// 时间复杂度O(nlogn + mlogn)
// 假设根节点为root输入m个在线查询，每个查询两个节点x, y，求x, y的最近公共祖先
func (t *Tree) DoubleLCA(root int, m int) {
	n := len(t.children)
	fa := make([][20]int, n)
	deep := make([]int, n)
	var dfs func(int, int)
	dfs = func(x, father int) {
		deep[x] = deep[father] + 1
		fa[x][0] = father
		for i := 1; 1<<i <= deep[x]; i++ {
			fa[x][i] = fa[fa[x][i-1]][i-1]
		}
		for _, y := range t.children[x] {
			if y.index == father {
				continue
			}
			dfs(y.index, x)
		}
	}
	dfs(root, 0)
	LCA := func(x, y int) int {
		if deep[x] < deep[y] {
			x, y = y, x
		}
		for i := 19; i >= 0; i-- {
			if deep[fa[x][i]]-1<<i >= deep[y] {
				x = fa[x][i]
			}
		}
		if x == y {
			return x
		}
		for i := 19; i >= 0; i-- {
			if fa[x][i] != fa[y][i] {
				x = fa[x][i]
				y = fa[y][i]
			}
		}
		return fa[x][0]
	}
	defer bufio.NewWriter(os.Stdout).Flush()
	for m > 0 {
		m--
		var x, y int
		// 输入x, y
		// 输出LCA(x, y)
		fmt.Fscan(bufio.NewReader(os.Stdin), &x, &y)
		fmt.Fprintln(bufio.NewWriter(os.Stdout), LCA(x, y))
	}
}

// Tarjan 算法求LCA，支持离线查询,时间复杂度O(n + m)， 核心思想bfs + 并查集
// 假设根节点为root输入m个离线查询，每个查询两个节点x, y，求x, y的最近公共祖先，返回m个查询的答案
func (t *Tree) Tarjan(root int, m int) []int {
	ans := make([]int, m)

	return ans
}
