package sliding_log

import (
	"sync"
	"time"
)

// Ограничиваем кол-во запросов в interval (например в одну сек) сохраняет в журнале время когда обращались к лимитеру
type SlidingLogLimiter struct {
	limit    int
	interval time.Duration
	logs     []time.Time
	mutex    sync.Mutex
}

func NewSlidingLogLimiter(limit int, interval time.Duration) *SlidingLogLimiter {
	return &SlidingLogLimiter{
		limit:    limit,
		interval: interval,
		logs:     make([]time.Time, 0),
	}
}

func (l *SlidingLogLimiter) allow() bool {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	lastReriod := time.Now().Add(-l.interval)
	for len(l.logs) != 0 && l.logs[0].Before(lastReriod) {
		l.logs = l.logs[1:]
	}
	l.logs = append(l.logs, time.Now())

	return len(l.logs) <= l.limit
}
