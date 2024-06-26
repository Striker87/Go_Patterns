package fixed_window

import (
	"context"
	"sync/atomic"
	"time"
)

// https://www.youtube.com/watch?v=w4suQQtnYmY
// Ограничиваем кол-во запросов в interval (например в одну сек)
type FixedWindowLimiter struct {
	count int32
	limit int32
}

func NewFixedWindowLimiter(ctx context.Context, limit int32, interval time.Duration) *FixedWindowLimiter {
	limiter := &FixedWindowLimiter{
		count: 0,
		limit: limit,
	}

	go limiter.startPeriodicCountRefresh(ctx, interval) // занюляем счетчик

	return limiter
}

func (l *FixedWindowLimiter) startPeriodicCountRefresh(ctx context.Context, interval time.Duration) {
	timer := time.NewTicker(interval)
	defer timer.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-timer.C:
			atomic.StoreInt32(&l.count, 0)
		}
	}
}

func (l *FixedWindowLimiter) Allow() bool { // проверяем исчерпан лимит запросов или нет
	count := atomic.LoadInt32(&l.count)
	if count > l.limit {
		return false
	}

	for !atomic.CompareAndSwapInt32(&l.count, count, count+1) {
		count = atomic.LoadInt32(&l.count)
	}

	return count < l.limit
}
