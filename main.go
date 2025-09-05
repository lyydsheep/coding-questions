package main

import (
	"coding-questions/algorithm"
	"fmt"
)

func main() {
	defer algorithm.Writer.Flush()
	var n int
	fmt.Fscan(algorithm.Reader, &n)
	fmt.Fprint(algorithm.Writer, n)
}
