package limit

import (
	"sync"
	"time"
)

type Counter struct {
	interval time.Duration
	maxCount int64
	count    int64
	lock     sync.Mutex
	ticker   *time.Ticker
}

func NewConter(maxCount int64, interval time.Duration) *Counter {
	c := &Counter{
		interval: interval,
		maxCount: maxCount,
		ticker: time.NewTicker(interval),
	}
	go func() {
		for {
			select {
			case <-c.ticker.C:
				c.lock.Lock()
				c.count = 0
				c.lock.Unlock()
			}
		}
	}()
	return c
}

// counter run .
func (c *Counter) Run() bool {
	c.lock.Lock()
	defer c.lock.Unlock()
	if c.count > c.maxCount {
		return false
	} else {
		c.count++
		return true
	}
}