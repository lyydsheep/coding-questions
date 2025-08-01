package questions

import (
	"context"
	"fmt"
	"sync"
)

// SolveQ5 用两个协程分别打印出一个数组中所有偶数的和、所有奇数的和
func SolveQ5(ctx context.Context) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	oddCh, evenCh := make(chan int), make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		sum := 0
		for i := range oddCh {
			sum += i
		}
		fmt.Println("odd:", sum)
	}()
	go func() {
		defer wg.Done()
		sum := 0
		for i := range evenCh {
			sum += i
		}
		fmt.Println("even:", sum)
	}()

	for _, x := range nums {
		if x%2 == 0 {
			evenCh <- x
			continue
		}
		oddCh <- x
	}
	close(oddCh)
	close(evenCh)
	wg.Wait()
}
