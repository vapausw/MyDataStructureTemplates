package main

//此文档用于记录有关于图的一些常用算法

const (
	inf int = 1e9 // 极大值
)

type node struct { // 存储图中的节点信息，存储当前的节点值，以及与其相连的边权值
	index, value int
}
type Graph struct {
	graph               [][]node // 存储图一般存储相连的便即可, 并且将自己的信息存储到第一个节点，例如1相连的边为2， 3， 4，那么g[1] = {2, 3, 4}
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
		u := q.Pop().(int)
		res = append(res, u)
	}
	if len(res) == n {
		return
	} else {
		return []int{}
	}
}
