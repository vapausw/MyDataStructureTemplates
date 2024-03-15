package main

import (
	"container/heap"
	"math"
)

//此文档用于记录有关于图的一些常用算法

const (
	inf int = math.MaxInt >> 1 // 极大值
)

type node struct { // 存储图中的节点信息，存储当前的节点值，以及与其相连的边权值
	index, value int
}
type Graph struct {
	graph               [][]node // 存储图一般存储相连的便即可，例如1相连的边为2， 3， 4，那么g[1] = {2, 3, 4}
	inDegree, outDegree []int    // 存储当前节点的入度和出度
}

func (g *Graph) New(n int) { // 创建一个有n个节点的图
	g.graph = make([][]node, n)
	g.inDegree = make([]int, n)
	g.outDegree = make([]int, n)
}

func (g *Graph) Directed(x, y, value int) { //添加有向边，从x指向y的边
	g.graph[x] = append(g.graph[x], node{y, value})
	g.outDegree[x]++
	g.inDegree[y]++
}

func (g *Graph) Ndirected(x, y, value int) { //添加无向边
	g.graph[x] = append(g.graph[x], node{y, value})
	g.graph[y] = append(g.graph[y], node{x, value})
	g.inDegree[x]++
	g.outDegree[x]++
	g.inDegree[y]++
	g.outDegree[y]++
}

func (g *Graph) Dfs(x node, vis []bool) { //递归遍历实现, vis记录当前节点是否被访问
	vis[x.index] = true
	for _, y := range g.graph[x.index] {
		if !vis[y.index] {
			g.Dfs(y, vis)
		}
	}
}

func (g *Graph) Bfs() {
	vis := make([]bool, len(g.graph))
	q := newQueue() // 此处使用queue.go中定义的队列
	q.Push(g.graph[0])
	vis[0] = true
	for !q.Empty() {
		x := q.Pop().(node)
		for _, y := range g.graph[x.index] {
			if vis[y.index] {
				continue
			}
			q.Push(y)
		}
	}
}

// Toposort 拓扑排序
func (g *Graph) Toposort() (res []int) { // 返回拓扑排序结果，如果为空即存在环，不存在拓扑排序的结果
	n := len(g.graph)
	in := make([]int, n) // 将每个节点的入度拿出来
	copy(in, g.inDegree)
	q := newQueue()
	for i := 0; i < n; i++ {
		if in[i] == 0 {
			q.Push(i)
		}
	}
	for !q.Empty() {
		x := q.Pop().(int)
		res = append(res, x)
		for _, y := range g.graph[x] {
			in[y.index]--
			if in[y.index] == 0 {
				q.Push(y.index)
			}
		}
	}
	if len(res) == n {
		return
	} else {
		return []int{}
	}
}

// DAG, dp求单源最短路以及最长路,时间复杂度O(n + m)

func (g *Graph) Dagdp(x int) (min_dis []int, max_dis []int) { //以x为起点的最短路和最长路
	topo := g.Toposort()
	min_dis = make([]int, len(g.graph))
	max_dis = make([]int, len(g.graph))
	for i := range min_dis {
		min_dis[i] = inf
	}
	min_dis[x] = 0
	for _, u := range topo {
		for _, y := range g.graph[u] {
			min_dis[y.index] = min(min_dis[y.index], min_dis[u]+y.value)
			max_dis[y.index] = max(max_dis[y.index], max_dis[u]+y.value)
		}
	}
	return min_dis, max_dis
}

// 最短路算法，floyd算法，时间复杂度O(n^3)

func (g *Graph) Floyd() [][]int {
	n := len(g.graph)
	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, n)
		for j := range f[i] {
			f[i][j] = inf
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < len(g.graph[i]); j++ {
			f[i][g.graph[i][j].index] = g.graph[i][j].value
			f[g.graph[i][j].index][i] = g.graph[i][j].value // 有向图将这行注释即可
		}
	}
	for k := 0; k < n; k++ {
		for x := 0; x < n; x++ {
			for y := 0; y < n; y++ {
				f[x][y] = min(f[x][y], f[x][k]+f[k][y])
			}
		}
	}
	return f
}

// 最短路算法，dijkstra算法，时间复杂度O(n^2 * log(n))，此处最短路时任意两点之间的最短路

func (gr *Graph) Dijkstra() (ans [][]int) {
	const inf int = 1e9 + 7
	n := len(gr.graph)
	dist := make([]int, n)
	g := make([][]pair, n)
	for i := range gr.graph {
		for j := range gr.graph[i] {
			g[i] = append(g[i], pair{gr.graph[i][j].index, gr.graph[i][j].value})
			g[gr.graph[i][j].index] = append(g[gr.graph[i][j].index], pair{i, gr.graph[i][j].value})
		}
	}
	dijkstra := func(u int) []int { //朴素的dijkstra
		for i := range dist {
			dist[i] = inf
		}
		h := hp{}
		heap.Push(&h, pair{u, 0})
		for h.Len() > 0 {
			p := heap.Pop(&h).(pair)
			if dist[p.i] < inf {
				continue
			}
			dist[p.i] = p.v
			for _, p1 := range g[p.i] {
				if dist[p1.i] == inf {
					heap.Push(&h, pair{p1.i, p1.v + p.v})
				}
			}
		}
		return dist
	}
	for i := 0; i < n; i++ {
		ans = append(ans, dijkstra(i))
	}
	return
}

type pair struct {
	i, v int
}

type hp []pair

func (h *hp) Len() int {
	return len(*h)
}

func (h *hp) Less(i, j int) bool {
	a := *h
	return a[i].v < a[j].v
}

func (h *hp) Swap(i, j int) {
	a := *h
	a[i], a[j] = a[j], a[i]
}

func (h *hp) Push(x any) {
	*h = append(*h, x.(pair))
}

func (h *hp) Pop() any {
	a := *h
	v := a[len(a)-1]
	*h = a[:len(a)-1]
	return v
}
