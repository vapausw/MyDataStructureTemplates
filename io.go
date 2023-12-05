package main

import (
	"bufio"
	. "fmt"
	"os"
)

// 此处记录输入输出的模板
// codeforces模板

func Cf_01_A() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	out.Flush()
	var t int
	for Fscan(in, t); t > 0; t-- {
		Fprintln(out, t) // 有换行输出
		Fprint(out, t)   // 无换行输出
	}
}
