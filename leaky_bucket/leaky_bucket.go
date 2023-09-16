package leaky_bucket

import (
	"context"
	"time"
)

// Ограничиваем кол-во запросов в interval (например в одну сек)
// забираем токены один раз в какой-то интервал
// каждый запрос добавляет токен
// если лимит токенов исчерпан, нет возможности пройти дальше (rate limiter), запрос отбрасывается

type LeakyBucketLimiter struct {
	leakyBucketCh chan struct{}
}

func NewLeakyBucketLimiter(ctx context.Context, limit int, period time.Duration) *LeakyBucketLimiter {
	limiter := &LeakyBucketLimiter{
		leakyBucketCh: make(chan struct{}),
	}

	leakInterval := period.Nanoseconds() / int64(limit)
	go limiter.startPeriodicLeak(ctx, time.Duration(leakInterval))

	return limiter
}

func (l *LeakyBucketLimiter) startPeriodicLeak(ctx context.Context, interval time.Duration) {
	timer := time.NewTicker(interval)
	defer timer.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-timer.C:
			select {
			case <-l.leakyBucketCh: // удаляем токен
			default:
			}
		}
	}
}

func (l *LeakyBucketLimiter) Allow() bool {
	select {
	case l.leakyBucketCh <- struct{}{}:
		return true
	default:
		return false
	}
}
