package questions

import (
	"context"
	"fmt"
	"time"
)

// Q1: 使用三个协程，每秒钟打印轮流打印 ABC
func SolveQ1(ctx context.Context) {
	ch1, ch2, ch3 := make(chan string), make(chan string), make(chan string)

	go worker(ctx, ch1, ch2, "B")

	go worker(ctx, ch2, ch3, "C")

	go worker(ctx, ch3, ch1, "A")

	ch1 <- "A"
}

func worker(ctx context.Context, ch1, ch2 chan string, ne string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("exit...")
			return
		case str := <-ch1:
			fmt.Println(str)
			time.Sleep(time.Second)
			select {
			case ch2 <- ne:
			case <-ctx.Done():
				fmt.Println("exit...")
				return
			}
		}
	}
}
