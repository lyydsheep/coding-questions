package tips

import "testing"

// 给切片预分配容量

func BenchmarkSliceWithoutPreAlloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var slice []int
		for j := 0; j < 10000; j++ {
			slice = append(slice, j)
		}
	}
}

func BenchmarkSliceWithPreAlloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slice := make([]int, 0, 10000)
		for j := 0; j < 10000; j++ {
			slice = append(slice, j)
		}
	}
}
