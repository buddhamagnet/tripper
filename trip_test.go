package tripper

import (
	"errors"
	"testing"
)

var failed, tripped int

func trip() {
	tripped++
}

func fail() {
	failed++
}

func TestTrip(t *testing.T) {
	cb := tripThreshold(10, trip, fail, func() {}, func() {}, func() {})
	for i := 0; i <= 100; i++ {
		cb.Call(func() error {
			return errors.New("failed")
		}, 0)
	}
	if failed != 10 {
		t.Errorf("expected 10 failures, got %d\n", failed)
	}
	if tripped != 1 {
		t.Errorf("expected 1 trip, got %d\n", tripped)
	}
}
