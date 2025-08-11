package pkg

import (
	"fmt"
	"go.uber.org/goleak"
	"testing"
	"time"
)

func TestLeak(t *testing.T) {
	go func() {
		time.Sleep(time.Second)
	}()
	fmt.Println("leak test")
}

func TestCommon(t *testing.T) {
	fmt.Println("common test")
}

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
}
