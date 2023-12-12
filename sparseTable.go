package main

import "math/bits"

// st表，调表解决RMQ问题，最大最小区间问题
// st[i][j] 对应的区间是 [i, i+2^j)
// 离线查询，不支持更改

type ST [][]int

func NewSt(a []int) ST {
	n := len(a)
	sz := bits.Len(uint(n)) // 二进制位数
	st := make(ST, n)
	for i, v := range a {
		st[i] = make([]int, sz)
		st[i][0] = v
	}
	for j := 1; 1<<j < n; j++ {
		for i := 0; i+1<<j <= n; i++ {
			st[i][j] = st.Op(st[i][j-1], st[i+1<<(j-1)][j-1])
		}
	}
	return st
}

func (st ST) Query(l, r int) int {
	k := bits.Len(uint(r-l)) - 1
	return st.Op(st[l][k], st[r-1<<k][k])
}

// Op 此处维护最大值，最小值，或者gcd
// 当前此处维护最大值
func (st ST) Op(a, b int) (res int) {
	res = max(a, b)
	return
}
