package clock

import (
	"time"
)

type Clock struct {
	startTime time.Time
}

func NewClock() Clock {
	return Clock{
		startTime: time.Now(),
	}
}

func (c *Clock) StopClock() time.Duration {
	return time.Since(c.startTime)
}
