package tips

import (
	"fmt"
	"testing"
	"time"
)

// 一行代码测量函数的执行时间

func TestTrackTime(t *testing.T) {
	defer TrackTime(time.Now())
	// do something
	time.Sleep(time.Second)
}

func TrackTime(pre time.Time) {
	fmt.Println(time.Since(pre))
}
