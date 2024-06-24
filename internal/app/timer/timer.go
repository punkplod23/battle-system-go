package timer

import (
	"time"
)

type Timer struct {
	Seconds     int
	Expired     bool
	CurrentTime int
}

func NewTimer(Seconds int) *Timer {
	t := new(Timer)
	t.Seconds = Seconds
	t.Expired = false
	t.CurrentTime = Seconds
	return t
}

func (t *Timer) StartCountdown() *Timer {
	for t.CurrentTime >= 0 {
		time.Sleep(1 * time.Second)
		t.CurrentTime = t.CurrentTime - 1
		if t.CurrentTime <= 0 {
			break
		}
	}
	return t
}

func (t *Timer) CheckComplete() bool {
	if t.CurrentTime <= 0 {
		return true
	}
	return false
}

func (t *Timer) ResetTimer() *Timer {
	t.CurrentTime = t.Seconds
	return t
}
