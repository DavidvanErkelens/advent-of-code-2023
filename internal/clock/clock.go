package clock

import (
	"fmt"
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

func (c *Clock) StopClock() {
	runtime := time.Now().Sub(c.startTime)
	fmt.Printf("Runtime: %dms (or %dμs)\n", runtime.Milliseconds(), runtime.Microseconds())
}
