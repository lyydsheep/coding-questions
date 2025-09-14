package main

import (
	"bufio"
	. "fmt"
	"os"
)

// ACM 输入输出
var (
	r *bufio.Reader
	w *bufio.Writer
)

func init() {
	r = bufio.NewReader(os.Stdin)
	w = bufio.NewWriter(os.Stdout)
}

func main() {
	defer w.Flush()
	var n, m int
	Fscan(r, &n, &m)
	// 滚动数组
	f, g := make([]int, m), make([]int, m+1)
	q := make([]int, n)

	for i := 0; i < n; i++ {
		var v, w, s int
		Fscan(r, &v, &w, &s)

		copy(g, f)
		// 枚举余数
		for j := 0; j < v; j++ {
			// 单调队列
			hh, tt := 0, -1
			for k := j; k <= m; k += v {
				// 移除队头
				if tt >= hh && q[hh] < k-s*v {
					hh++
				}

				// 取出窗口内最大值
				f[k] = max(f[k], g[q[hh]]+(k-q[hh])/v*w)

				// 加入窗口
				for tt >= hh && g[q[tt]]-(q[tt]-j)/v*w <= f[k]-(k-j)/v*w {
					tt--
				}
				tt++
				q[tt] = k
			}
		}
	}

	Fprintln(w, f[m])
}
