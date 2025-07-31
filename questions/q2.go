package questions

import (
	"context"
	"fmt"
	"sync"
)

// SolveQ2 : 实现两个协程轮流输出 A1B2C3 ... Z26
func SolveQ2(ctx context.Context) {
	// 不考虑优雅退出
	ch1, ch2 := make(chan rune), make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := range ch1 {
			fmt.Printf("%c\n", i)
			ch2 <- int(i - 'A' + 1)
		}
	}()
	go func() {
		defer wg.Done()
		for i := range ch2 {
			fmt.Println(i)
			if i == 26 {
				close(ch1)
				return
			}
			ch1 <- rune('A' + i)
		}
	}()

	ch1 <- 'A'
	wg.Wait()
	close(ch2)
}
