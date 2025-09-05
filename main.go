package main

import (
	"bufio"
	. "fmt"
	"os"
)

var (
	w *bufio.Writer
	r *bufio.Reader
)

func init() {
	w, r = bufio.NewWriter(os.Stdout), bufio.NewReader(os.Stdin)
}

func solve(a, b []int) {
	n := len(a)
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}

	for i := 1; i <= n; i++ {
		maxv := 0
		for j := 1; j <= n; j++ {
			// 没有 a[i]
			f[i][j] = f[i-1][j]

			// 有 a[i]
			if a[i-1] == b[j-1] {
				f[i][j] = max(f[i][j], maxv)
			}

			maxv = max(maxv+1, f[i-1][j])
		}
	}
}

func main() {
	defer w.Flush()
	var n int
	Fscan(r, &n)

	a, b := make([]int, n), make([]int, n)
	for i := range a {
		Fscan(r, &a[i])
	}
	for i := range b {
		Fscan(r, &b[i])
	}

	solve(a, b)
}
