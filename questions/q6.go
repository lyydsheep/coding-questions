package questions

import (
	"context"
	"fmt"
	"math/rand"
)

// SolveQ6 多个协程执行随机数加法，最后输出其中的最大值
func SolveQ6(ctx context.Context) {
	n := 10
	res := make(chan int)

	for i := range n {
		go func() {
			sum := rand.Intn(100) + rand.Intn(100)
			fmt.Println(i, sum)
			res <- sum
		}()
	}

	ans := -1
	for range n {
		if num := <-res; num > ans {
			ans = num
		}
	}
	fmt.Println("ans", ans)
}
