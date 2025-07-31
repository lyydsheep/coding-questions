package questions

import (
	"context"
	"fmt"
	"runtime"
	"sync"
)

// SolveQ4 给出代码输出，重点考察 gmp 协程调度过程
func SolveQ4(ctx context.Context) {
	runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup
	n := 10
	wg.Add(n)

	for i := range n {
		go func() {
			defer wg.Done()
			fmt.Println("i am goroutine", i)
		}()
	}

	wg.Wait()
}
