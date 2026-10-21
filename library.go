package pas2

import (
	"fmt"
	"math"
	"sync"
	"time"
)

type Bucket struct {
	mu                          sync.Mutex
	capacity, tokens, perSecond float64
	last                        time.Time
}

func New(capacity, perSecond float64, now time.Time) (*Bucket, error) {
	if capacity < 1 || perSecond <= 0 || math.IsNaN(capacity) || math.IsInf(capacity, 0) || math.IsNaN(perSecond) || math.IsInf(perSecond, 0) {
		return nil, fmt.Errorf("invalid capacity or rate")
	}
	return &Bucket{capacity: capacity, tokens: capacity, perSecond: perSecond, last: now}, nil
}
func (b *Bucket) Take(now time.Time, cost float64) bool {
	b.mu.Lock()
	defer b.mu.Unlock()
	if cost <= 0 || math.IsNaN(cost) || math.IsInf(cost, 0) || cost > b.capacity {
		return false
	}
	if now.After(b.last) {
		b.tokens = math.Min(b.capacity, b.tokens+now.Sub(b.last).Seconds()*b.perSecond)
		b.last = now
	}
	if cost > b.tokens {
		return false
	}
	b.tokens -= cost
	return true
}
