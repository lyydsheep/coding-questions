package pkg

import (
	"github.com/panjf2000/ants/v2"
	"golang.org/x/sync/errgroup"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

const (
	RunTimes           = 1e6
	PoolCap            = 5e4
	BenchParam         = 15
	DefaultExpiredTime = 10 * time.Second
)

// 计算密集型任务：计算斐波那契数列，使用递归方式实现，会消耗较多 CPU 资源
func demoFunc() {
	var fib func(n int) int
	fib = func(n int) int {
		if n <= 1 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}
	// 计算一个较大的斐波那契数，确保是计算密集型操作
	_ = fib(BenchParam)
}

func demoPoolFunc(args any) {
	n := args.(int)
	time.Sleep(time.Duration(n) * time.Millisecond)
}

func demoPoolFuncInt(n int) {
	time.Sleep(time.Duration(n) * time.Millisecond)
}

var stopLongRunningFunc int32

func longRunningFunc() {
	for atomic.LoadInt32(&stopLongRunningFunc) == 0 {
		runtime.Gosched()
	}
}

func longRunningPoolFunc(arg any) {
	<-arg.(chan struct{})
}

func longRunningPoolFuncCh(ch chan struct{}) {
	<-ch
}

func BenchmarkGoroutines(b *testing.B) {
	var wg sync.WaitGroup
	for i := 0; i < b.N; i++ {
		wg.Add(RunTimes)
		for j := 0; j < RunTimes; j++ {
			go func() {
				demoFunc()
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

func BenchmarkChannel(b *testing.B) {
	var wg sync.WaitGroup
	sema := make(chan struct{}, PoolCap)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wg.Add(RunTimes)
		for j := 0; j < RunTimes; j++ {
			sema <- struct{}{}
			go func() {
				demoFunc()
				<-sema
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

func BenchmarkErrGroup(b *testing.B) {
	var wg sync.WaitGroup
	var pool errgroup.Group
	pool.SetLimit(PoolCap)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wg.Add(RunTimes)
		for j := 0; j < RunTimes; j++ {
			pool.Go(func() error {
				demoFunc()
				wg.Done()
				return nil
			})
		}
		wg.Wait()
	}
}

func BenchmarkAntsPool(b *testing.B) {
	var wg sync.WaitGroup
	p, _ := ants.NewPool(PoolCap, ants.WithExpiryDuration(DefaultExpiredTime))
	defer p.Release()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wg.Add(RunTimes)
		for j := 0; j < RunTimes; j++ {
			_ = p.Submit(func() {
				demoFunc()
				wg.Done()
			})
		}
		wg.Wait()
	}
}

func BenchmarkAntsMultiPool(b *testing.B) {
	var wg sync.WaitGroup
	p, _ := ants.NewMultiPool(10, PoolCap/10, ants.RoundRobin, ants.WithExpiryDuration(DefaultExpiredTime))
	defer p.ReleaseTimeout(DefaultExpiredTime) //nolint:errcheck

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wg.Add(RunTimes)
		for j := 0; j < RunTimes; j++ {
			_ = p.Submit(func() {
				demoFunc()
				wg.Done()
			})
		}
		wg.Wait()
	}
}

func BenchmarkGoroutinesThroughput(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < RunTimes; j++ {
			go demoFunc()
		}
	}
}

func BenchmarkSemaphoreThroughput(b *testing.B) {
	sema := make(chan struct{}, PoolCap)
	for i := 0; i < b.N; i++ {
		for j := 0; j < RunTimes; j++ {
			sema <- struct{}{}
			go func() {
				demoFunc()
				<-sema
			}()
		}
	}
}

func BenchmarkAntsPoolThroughput(b *testing.B) {
	p, _ := ants.NewPool(PoolCap, ants.WithExpiryDuration(DefaultExpiredTime))
	defer p.Release()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < RunTimes; j++ {
			_ = p.Submit(demoFunc)
		}
	}
}

func BenchmarkAntsMultiPoolThroughput(b *testing.B) {
	p, _ := ants.NewMultiPool(10, PoolCap/10, ants.RoundRobin, ants.WithExpiryDuration(DefaultExpiredTime))
	defer p.ReleaseTimeout(DefaultExpiredTime) //nolint:errcheck

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < RunTimes; j++ {
			_ = p.Submit(demoFunc)
		}
	}
}

func BenchmarkParallelAntsPoolThroughput(b *testing.B) {
	p, _ := ants.NewPool(PoolCap, ants.WithExpiryDuration(DefaultExpiredTime))
	defer p.Release()

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = p.Submit(demoFunc)
		}
	})
}

func BenchmarkParallelAntsMultiPoolThroughput(b *testing.B) {
	p, _ := ants.NewMultiPool(10, PoolCap/10, ants.RoundRobin, ants.WithExpiryDuration(DefaultExpiredTime))
	defer p.ReleaseTimeout(DefaultExpiredTime) //nolint:errcheck

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = p.Submit(demoFunc)
		}
	})
}

/** 测试结果

goos: darwin
goarch: arm64
pkg: coding-questions/questions
cpu: Apple M4 Pro
BenchmarkGoroutines-12                                 4         289920823 ns/op        16000208 B/op    1000002 allocs/op
BenchmarkChannel-12                                    2         768443625 ns/op        24000180 B/op    1000003 allocs/op
BenchmarkErrGroup-12                                   2         848165125 ns/op        40380900 B/op    2000791 allocs/op
BenchmarkAntsPool-12                                   2         900733354 ns/op        16130420 B/op    1001546 allocs/op
BenchmarkAntsMultiPool-12                              2         520914584 ns/op        16258876 B/op    1003134 allocs/op
BenchmarkGoroutinesThroughput-12                       4         291706115 ns/op              12 B/op          0 allocs/op
BenchmarkSemaphoreThroughput-12                        2         751605479 ns/op        16000092 B/op    1000002 allocs/op
BenchmarkAntsPoolThroughput-12                         2         858998688 ns/op           66860 B/op        968 allocs/op
BenchmarkAntsMultiPoolThroughput-12                    3         480316014 ns/op          125274 B/op       1751 allocs/op
BenchmarkParallelAntsPoolThroughput-12           1592536               763.1 ns/op             0 B/op          0 allocs/op
BenchmarkParallelAntsMultiPoolThroughput-12      3182235               371.2 ns/op             0 B/op          0 allocs/op

*/
