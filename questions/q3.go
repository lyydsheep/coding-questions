package questions

import (
	"context"
	"fmt"
)

// SolveQ3 给出输出结果
func SolveQ3(ctx context.Context) {
	doAppend := func(s []int) {
		s = append(s, 1)
		printLenAndCap(s)
	}

	s := make([]int, 8)
	doAppend(s[:4])
	printLenAndCap(s)

	doAppend(s)
	printLenAndCap(s)
}

func printLenAndCap(s []int) {
	fmt.Println(s)
	fmt.Printf("len: %d, cap: %d\n", len(s), cap(s))
}
