package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f := newQueue()
	f.Push(1)
	v := f.Pop()
	_, err := fmt.Fprintln(bufio.NewWriter(os.Stdout), v)
	if err != nil {
		return
	}
}
