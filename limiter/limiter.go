package limiter

import (
	"sync"
	"sync/atomic"
)

type Limiter struct {
	num   int64
	ch    chan struct{}
	wg    sync.WaitGroup
	count int64
}

func NewLimiter(num int64) *Limiter {
	l := &Limiter{
		num: num,
		ch:  make(chan struct{}, num),
	}
	return l
}

func (l *Limiter) Allow() {
	l.ch <- struct{}{}
	l.wg.Add(1)
}

func (l *Limiter) Done() {
	<-l.ch
	l.wg.Done()
	atomic.AddInt64(&l.count, 1)
}

func (l *Limiter) Left() int64 {
	return l.num - int64(len(l.ch))
}

func (l *Limiter) Wait() {
	l.wg.Wait()
}

func (l *Limiter) Count() int64 {
	return l.count
}
