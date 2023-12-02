package main

import (
	"bufio"
	. "fmt"
	"os"
)

func main() {
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	Fprintln(out, "test")
}
