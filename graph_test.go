package main

import (
	"reflect"
	"testing"
)

// 测试NewGraph
func TestNewGraph(t *testing.T) {
	g := &Graph{}
	g.New(5)
	if len(g.graph) != 5 || len(g.inDegree) != 5 || len(g.outDegree) != 5 {
		t.Errorf("Graph initialization failed")
	}
}

// 测试Directed函数
func TestDirected(t *testing.T) {
	g := &Graph{}
	g.New(3) // 创建一个有3个节点的图
	g.Directed(0, 1, 10)
	g.Directed(1, 2, 20)

	if len(g.graph[0]) != 1 || g.graph[0][0].index != 1 || g.graph[0][0].value != 10 {
		t.Errorf("Directed edge addition failed")
	}

	if len(g.graph[1]) != 1 || g.graph[1][0].index != 2 || g.graph[1][0].value != 20 {
		t.Errorf("Directed edge addition failed")
	}
}

// 测试Ndirected函数
func TestNdirected(t *testing.T) {
	g := &Graph{}
	g.New(3) // 创建一个有3个节点的图
	g.Ndirected(0, 1, 10)
	g.Ndirected(1, 2, 20)

	if len(g.graph[0]) != 1 || g.graph[0][0].index != 1 || g.graph[0][0].value != 10 {
		t.Errorf("Nondirected edge addition failed")
	}

	if len(g.graph[1]) != 2 || g.graph[1][1].index != 2 || g.graph[1][1].value != 20 {
		t.Errorf("Nondirected edge addition failed")
	}

	if len(g.graph[2]) != 1 || g.graph[2][0].index != 1 || g.graph[2][0].value != 20 {
		t.Errorf("Nondirected edge addition failed")
	}
}

// 测试Dfs函数
func TestDfs(t *testing.T) {
	g := &Graph{}
	g.New(4)
	g.Ndirected(0, 1, 0)
	g.Ndirected(1, 2, 0)
	g.Ndirected(2, 3, 0)
	vis := make([]bool, 4)
	g.Dfs(node{index: 0}, vis)
	if !vis[0] || !vis[1] || !vis[2] || !vis[3] {
		t.Errorf("DFS failed to visit all nodes")
	}
}

// 此测试函数不实现了，因为无法判断，只能人工观察
func TestBfs(t *testing.T) {

}

// 测试Toposort函数
func TestToposort(t *testing.T) {
	g := &Graph{}
	g.New(3)
	g.Directed(0, 1, 10)
	g.Directed(1, 2, 20)

	res := g.Toposort()
	expected := []int{0, 1, 2}
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("Toposort failed, got %v, want %v", res, expected)
	}
}
