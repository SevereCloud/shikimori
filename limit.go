package shikimori

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type limit struct {
	maxRequest int
	duration   time.Duration

	lastTime time.Time
	requests int
	mux      sync.Mutex
}

func newLimit(maxRequest int, duration time.Duration) *limit {
	return &limit{
		maxRequest: maxRequest,
		duration:   duration,
		lastTime:   time.Time{},
		requests:   0,
		mux:        sync.Mutex{},
	}
}

func (l *limit) reset() {
	l.lastTime = time.Now()
	l.requests = 0
}

func (l *limit) RateLimit(ctx context.Context) error {
	l.mux.Lock()
	defer l.mux.Unlock()

	sleepTime := l.duration - time.Since(l.lastTime)
	if sleepTime < 0 {
		l.reset()
	}

	l.requests++

	if l.requests <= l.maxRequest {
		return nil
	}

	select {
	case <-time.After(sleepTime):
		l.reset()
	case <-ctx.Done():
		return fmt.Errorf("shikimori: %w", ctx.Err())
	}

	return nil
}

type limits struct {
	limits []*limit
}

func newLimits(args ...*limit) *limits {
	return &limits{
		limits: args,
	}
}

func (l *limits) RateLimit(ctx context.Context) error {
	for i := 0; i < len(l.limits); i++ {
		err := l.limits[i].RateLimit(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}
