package internal

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Limit struct {
	maxRequest int
	duration   time.Duration

	requests *fifo[time.Time]
	mux      sync.Mutex
}

func NewLimit(maxRequest int, duration time.Duration) *Limit {
	queue := newFifo[time.Time]()
	queue.Limit(maxRequest)

	return &Limit{
		maxRequest: maxRequest,
		duration:   duration,
		requests:   queue,
		mux:        sync.Mutex{},
	}
}

func (l *Limit) RateLimit(ctx context.Context) error {
	l.mux.Lock()
	defer l.mux.Unlock()

	frontRequest := l.requests.Front()
	if frontRequest == nil || l.requests.Len() < l.maxRequest {
		l.requests.Push(time.Now())

		return nil
	}

	sleepTime := l.duration - time.Since(*frontRequest)

	select {
	case <-time.After(sleepTime):
		l.requests.Push(time.Now())
	case <-ctx.Done():
		return fmt.Errorf("limit: %w", ctx.Err())
	}

	return nil
}

type Limits struct {
	limits []*Limit
}

func NewLimits(args ...*Limit) *Limits {
	return &Limits{
		limits: args,
	}
}

func (l *Limits) RateLimit(ctx context.Context) error {
	for i := 0; i < len(l.limits); i++ {
		err := l.limits[i].RateLimit(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}
