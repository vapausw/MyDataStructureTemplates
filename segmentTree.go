package main

type seg []struct {
	l, r int // 左右子区间
	val  int // 线段树的值
}

// 区间的维护，例如 + * | & ^ min max gcd等等
func (t seg) mergeInfo(a, b int) int {
	return min(a, b)
}

func (t seg) set(o, val int) {
	t[o].val = t.mergeInfo(t[o].val, val)
}

// 线段树的维护
func (t seg) maintain(o int) {
	t[o].val = t.mergeInfo(t[o<<1].val, t[o<<1|1].val)
}

// 线段树的建立
func (t seg) build(a []int, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t[o].val = a[l-1]
		return
	}
	m := l + (r-l)>>1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t.maintain(o)
}

// 创建一个新的线段树，a 的下标从 0 开始
func newSegmentTree(a []int) seg {
	n := len(a)
	if n == 0 {
		panic("slice can't be empty")
	}
	t := make(seg, 4*n+1)
	t.build(a, 1, 1, n)
	return t
}

// o=1  1<=i<=n, 线段树的单点更新，不带懒惰标记处没有实现区间更新，区间更新使用下述带有懒惰标记的线段树
func (t seg) update(o, i, val int) {
	if t[o].l == t[o].r {
		t.set(o, val)
		return
	}
	m := (t[o].l + t[o].r) >> 1
	if i <= m {
		t.update(o<<1, i, val)
	} else {
		t.update(o<<1|1, i, val)
	}
	t.maintain(o)
}

// o=1  [l,r] 1<=l<=r<=n， 线段树的区间查询
func (t seg) query(o, l, r int) int {
	if l <= t[o].l && t[o].r <= r {
		return t[o].val
	}
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if m < l {
		return t.query(o<<1|1, l, r)
	}
	vl := t.query(o<<1, l, r)
	vr := t.query(o<<1|1, l, r)
	return t.mergeInfo(vl, vr)
}

func (t seg) queryAll() int { return t[1].val }

// 线段树上类似于lower_bound之类的东西
// 返回整个区间小于v的最靠做的位置
// 此时线段树维护的是区间最小值，其他的维护类似
func (t seg) queryFirstLessPos(o, v int) int {
	if t[o].l == t[o].r {
		return t[o].l
	}
	if t[o<<1].val < v {
		return t.queryFirstLessPos(o<<1, v)
	}
	return t.queryFirstLessPos(o<<1|1, v)
}

// 查询 [l,r] 上小于 v 的最靠左的位置
// 这里线段树维护的是区间最小值
// 不存在时返回 0,因为线段树的维护的区间是1 <= l <= r <= n
func (t seg) queryFirstLessPosInRange(o, l, r, v int) int {
	if t[o].val >= v {
		return 0
	}
	if t[o].l == t[o].r {
		return t[o].l
	}
	m := (t[o].l + t[o].r) >> 1
	if l <= m {
		if pos := t.queryFirstLessPosInRange(o<<1, l, r, v); pos > 0 {
			return pos
		}
	}
	if m < r {
		if pos := t.queryFirstLessPosInRange(o<<1|1, l, r, v); pos > 0 { // 注：这里 pos > 0 的判断可以省略，因为 pos == 0 时最后仍然会返回 0
			return pos
		}
	}
	return 0
}

const todoInit = 0

// 附带懒惰标记的线段树
type lazySeg []struct {
	l, r int
	val  int // 线段树保存的数据
	todo int // 懒惰标记保存的数据
}

// 维护+ * | & ^ min max gcd等等
func (t lazySeg) mergeInfo(a, b int) int {
	return a + b
}

func (t lazySeg) maintain(o int) {
	t[o].val = t.mergeInfo(t[o<<1].val, t[o<<1|1].val)
}

func (t lazySeg) do(o, val int) {
	to := &t[o]

	// 更新当前点对于整个区间的影响
	to.val += val * (to.r - to.l + 1)

	// 更新当前点对其左右儿子的影响
	to.todo += val
}

// 将懒惰标记向下堆
func (t lazySeg) spread(o int) {
	if v := t[o].todo; v != todoInit {
		t.do(o<<1, v)
		t.do(o<<1|1, v)
		t[o].todo = todoInit
	}
}

func (t lazySeg) build(a []int, o, l, r int) {
	t[o].l, t[o].r = l, r
	t[o].todo = todoInit
	if l == r {
		t[o].val = a[l-1]
		return
	}
	m := l + (r-l)>>1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t.maintain(o)
}

// a 的下标从 0 开始
func newLazySegmentTree(a []int) lazySeg {
	n := len(a)
	if n == 0 {
		panic("slice can't be empty")
	}
	t := make(lazySeg, 4*n+1)
	t.build(a, 1, 1, n)
	return t
}

// EXTRA: 适用于需要提取所有元素值的场景
func (t lazySeg) spreadAll(o int) {
	if t[o].l == t[o].r {
		return
	}
	t.spread(o)
	t.spreadAll(o << 1)
	t.spreadAll(o<<1 | 1)
}

func (t lazySeg) update(o, l, r, val int) {
	if l <= t[o].l && t[o].r <= r {
		t.do(o, val)
		return
	}
	t.spread(o)
	m := t[o].l + (t[o].r-t[o].l)>>1
	if l <= m {
		t.update(o<<1, l, r, val)
	}
	if m < r {
		t.update(o<<1|1, l, r, val)
	}
	t.maintain(o)
}

func (t lazySeg) query(o, l, r int) int {
	if l <= t[o].l && t[o].r <= r {
		return t[o].val
	}
	t.spread(o)
	m := t[o].l + (t[o].r-t[o].l)>>1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if l > m {
		return t.query(o<<1|1, l, r)
	}
	vl := t.query(o<<1, l, r)
	vr := t.query(o<<1|1, l, r)
	return t.mergeInfo(vl, vr)
}

func (t lazySeg) queryAll() int { return t[1].val }
