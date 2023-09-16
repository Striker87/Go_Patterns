package token_bucket

import (
	"context"
	"time"
)

// Ограничиваем кол-во запросов в interval (например в одну сек)
// есть пполнитель, который например каждые 100 мс добавляет токены
// если ведро переполнено, токен уходит
// когда приходит запрос, он забирает токен
// если лимит токенов исчерпан, нет возможности пройти дальше (rate limiter), запрос отбрасывается
type TokenBucketLimiter struct {
	tokeBucketCh chan struct{}
}

func NewTokenBucketLimiter(ctx context.Context, limit int, period time.Duration) *TokenBucketLimiter {
	limiter := &TokenBucketLimiter{
		tokeBucketCh: make(chan struct{}, limit),
	}

	for i := 0; i < limit; i++ {
		limiter.tokeBucketCh <- struct{}{}
	}

	replenishmentInterval := period.Nanoseconds() / int64(limit)

	go limiter.startPeriodicReplenishment(ctx, time.Duration(replenishmentInterval))
}

func (l *TokenBucketLimiter) startPeriodicReplenishment(ctx context.Context, interval time.Duration) {
	timer := time.NewTicker(interval)
	defer timer.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-timer.C:
			select {
			case l.tokeBucketCh <- struct{}{}: // добавляем токен
			default: // токен не будет добавлен
			}
		}
	}
}

func (l *TokenBucketLimiter) Allow() bool {
	select {
	case <-l.tokeBucketCh:
		return true
	default:
		return false
	}
}
