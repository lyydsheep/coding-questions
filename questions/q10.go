package questions

import (
	"sync"
	"time"
)

// LeakBucket 漏桶算法实现
type LeakBucket struct {
	capacity int            // 桶的容量
	rate     time.Duration  // 漏水速率（处理请求的速率）
	requests chan struct{}  // 请求队列
	quit     chan struct{}  // 退出信号
	wg       sync.WaitGroup // 等待组，用于优雅关闭
}

// NewLeakBucket 创建一个新的漏桶实例
// ratePerSecond: 每秒处理的请求数
// capacity: 桶的容量
func NewLeakBucket(ratePerSecond int, capacity int) *LeakBucket {
	b := &LeakBucket{
		capacity: capacity,
		rate:     time.Second / time.Duration(ratePerSecond),
		requests: make(chan struct{}, capacity),
		quit:     make(chan struct{}),
	}
	b.wg.Add(1)
	go b.startLeaking()
	return b
}

// startLeaking 定期处理请求的协程（漏水）
func (b *LeakBucket) startLeaking() {
	defer b.wg.Done()
	ticker := time.NewTicker(b.rate)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// 尝试处理一个请求（漏水）
			select {
			case <-b.requests:
				// 成功处理一个请求
			default:
				// 没有等待处理的请求
			}
		case <-b.quit:
			// 处理所有剩余请求
			for len(b.requests) > 0 {
				<-b.requests
			}
			return
		}
	}
}

// Wait 等待直到请求被处理（进入队列并等待处理）
func (b *LeakBucket) Wait() bool {
	select {
	case b.requests <- struct{}{}:
		// 请求成功进入队列
		return true
	default:
		// 队列已满，拒绝请求
		return false
	}
}

// Process 处理一个请求，结合漏桶限流
func (b *LeakBucket) Process(requestID int, handler func(int)) bool {
	// 尝试将请求加入处理队列
	if !b.Wait() {
		// 桶满了，拒绝请求
		return false
	}

	// 请求已加入队列，将在后台按固定速率处理
	// 注意：这里不直接调用handler，而是让漏桶按固定速率处理
	// 在实际应用中，handler应该在startLeaking中调用
	return true
}

// Close 关闭漏桶
func (b *LeakBucket) Close() {
	close(b.quit)
	b.wg.Wait()
}
