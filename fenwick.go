package main

// 树状数组，支持单点修改
type fenwick []int

// 防止死循环从1开始

func NewFenwick(n int) fenwick {
	f := make([]int, n+1)
	return f
}

func (f fenwick) Add(i, val int) {
	for ; i < len(f); i += i & -i {
		f[i] += val
	}
}

func (f fenwick) pre(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res += f[i]
	}
	return
}

func (f fenwick) Query(l, r int) int {
	return f.pre(r) - f.pre(l-1)
}

// 二维树状数组，支持区间修改,利用差分树状数组
// 利用差分数组，实现 O(log n) 的区间加、区间查询
// a[1] = diff[1]
// a[2] = diff[1] + diff[2]
// a[m] = diff[1] + ... + diff[m]
// 所以 a[1] + ... + a[m]
//
//	= ∑(m-i+1)*diff[i]
//	= (m+1)∑diff[i] - ∑i*diff[i]
type fenwickDiff [][2]int

func newFenwickDiff(n int) fenwickDiff {
	return make([][2]int, n+1)
}

// 单点修改
func (f fenwickDiff) _add(i, val int) {
	for iv := i * val; i < len(f); i += i & -i {
		f[i][0] += val
		f[i][1] += iv
	}
}

func (f fenwickDiff) add(l, r, val int) {
	f._add(l, val)
	f._add(r+1, -val)
}

func (f fenwickDiff) pre(i0 int) int {
	var s0, s1 int
	for i := i0; i > 0; i &= i - 1 {
		s0 += f[i][0]
		s1 += f[i][1]
	}
	return (i0+1)*s0 - s1
}

func (f fenwickDiff) Query(l, r int) int {
	return f.pre(r) - f.pre(l-1)
}
