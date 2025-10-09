package tips

// 通过 defer 实现在一个函数的开头、结尾做一些工作

import (
	"fmt"
	"testing"
	"time"
)

func TestMultiStageDefer(t *testing.T) {
	defer MultiStageDefer()()
	fmt.Println("do something")
	time.Sleep(time.Second)
}

func MultiStageDefer() func() {
	fmt.Println("do something init")

	return func() {
		fmt.Println("do something end")
	}
}
