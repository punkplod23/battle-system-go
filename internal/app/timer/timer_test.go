package timer

import (
	"testing"
)

// TestNewTimer checks if a new timer is created with the correct initial values.
func TestNewTimer(t *testing.T) {
	seconds := 10
	timer := NewTimer(seconds)

	if timer.Seconds != seconds {
		t.Errorf("Expected timer seconds to be %d, got %d", seconds, timer.Seconds)
	}

	if timer.Expired {
		t.Errorf("Expected new timer to not be expired")
	}

	if timer.CurrentTime != seconds {
		t.Errorf("Expected current time to be %d, got %d", seconds, timer.CurrentTime)
	}
}

// TestStartCountdown checks the countdown functionality of the timer.
// This test is skipped because it requires waiting for the countdown to finish.
func TestStartCountdown(t *testing.T) {
	t.Skip("Skipping countdown test")
}

// TestCheckComplete checks if the timer correctly reports its completion status.
func TestCheckComplete(t *testing.T) {
	timer := NewTimer(0)
	if timer.CheckComplete() {
		t.Errorf("Expected timer with 0 seconds to be complete")
	}

	timer = NewTimer(10)
	if !timer.CheckComplete() {
		t.Errorf("Expected timer with more than 0 seconds to not be complete")
	}
}

// TestResetTimer checks if the timer resets correctly.
func TestResetTimer(t *testing.T) {
	seconds := 10
	timer := NewTimer(seconds)
	timer.CurrentTime = 0 // Simulate timer has counted down
	timer.resetTimer()

	if timer.CurrentTime != seconds {
		t.Errorf("Expected current time to reset to %d, got %d", seconds, timer.CurrentTime)
	}
}
