package tripper

import (
	"errors"
	"testing"
)

var failed, tripped, toggle int

var callbacks = map[string]func(){
	"trip": func() { tripped++ },
	"fail": func() { failed++ },
}

func reset() {
	failed = 0
	tripped = 0
}

func TestTripConsecutive(t *testing.T) {
	cb := NewBreaker(10, "consecutive", callbacks)
	for i := 0; i <= 10; i++ {
		cb.Call(errorOut, 0)
	}
	if failed != 10 {
		t.Errorf("expected 10 failures, got %d\n", failed)
	}
	if tripped != 1 {
		t.Errorf("expected 1 trip, got %d\n", tripped)
	}
	reset()
}

func TestPassConsecutive(t *testing.T) {
	cb := NewBreaker(10, "consecutive", callbacks)
	for i := 0; i <= 10; i++ {
		cb.Call(allGood, 0)
	}
	if failed != 0 {
		t.Errorf("expected 0 failures, got %d\n", failed)
	}
	if tripped != 0 {
		t.Errorf("expected 0 trips, got %d\n", tripped)
	}
	reset()
}

func errorOut() error {
	return errors.New("foobar")
}

func allGood() error {
	return nil
}
