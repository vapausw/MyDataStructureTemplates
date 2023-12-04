package main

type UnionFind struct {
	Fa     []int //每个节点的父节点坐标，初始化时将节点的父节点初始化为自己
	Groups int   // 连通分量的个数
}

func NewUnionFind(n int) UnionFind {
	fa := make([]int, n)
	for i := range fa {
		fa[i] = -1
	}
	return UnionFind{
		Fa:     fa,
		Groups: n,
	}
}

func (u *UnionFind) Find(x int) int { // 路径压缩寻找父节点
	if u.Fa[x] != x {
		u.Fa[x] = u.Find(u.Fa[x])
	}
	return u.Fa[x]
}

func (u *UnionFind) Merge(from, to int) int { // 合并还可采取启发式合并，深度小的合并到深度较大的节点
	x, y := u.Find(from), u.Find(to)
	if x == y {
		return -1
	}
	u.Fa[x] = y
	u.Groups--
	return y
}

func (u *UnionFind) Same(x, y int) bool {
	return u.Find(x) == u.Find(y)
}
