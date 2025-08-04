package questions

import (
	"sync"
	"sync/atomic"
)

type Task struct {
	f func() error
}

func NewTask(f func() error) *Task {
	return &Task{f: f}
}

type Pool struct {
	ch  chan *Task
	cnt int64
	cap int64
	sync.Mutex
}

func NewPool(ch chan *Task, cap int64) *Pool {
	return &Pool{
		ch:  ch,
		cap: cap,
	}
}

func (p *Pool) GetCnt() int64 {
	return atomic.LoadInt64(&p.cnt)
}

func (p *Pool) IncCnt() {
	atomic.AddInt64(&p.cnt, 1)
}

func (p *Pool) DecCnt() {
	atomic.AddInt64(&p.cnt, -1)
}

func (p *Pool) run() {
	p.IncCnt()
	go func() {
		defer p.DecCnt()
		for task := range p.ch {
			task.f()
		}
	}()
}

func (p *Pool) AddTask(t *Task) {
	p.Lock()
	defer p.Unlock()

	if p.GetCnt() < p.cap {
		// 可以启动一个新的协程
		p.run()
	}
	p.ch <- t
}
