package main

import (
	"coding-questions/questions"
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)

	// working...
	questions.SolveQ6(ctx)

	// 优雅退出
	<-sigChan
	cancel()
	// 等待一段时间确保协程退出
	fmt.Println("quiting....")
	time.Sleep(time.Second * 5)
}
