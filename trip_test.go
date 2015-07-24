package tripper

import (
	"errors"
	"testing"
)

var failed, tripped int

var callbacks = map[string]func(){
	"trip": func() { tripped++ },
	"fail": func() { failed++ },
}

func TestTrip(t *testing.T) {
	cb := tripThreshold(10, callbacks)
	for i := 0; i <= 100; i++ {
		cb.Call(errorOut, 0)
	}
	if failed != 10 {
		t.Errorf("expected 10 failures, got %d\n", failed)
	}
	if tripped != 1 {
		t.Errorf("expected 1 trip, got %d\n", tripped)
	}
}

func errorOut() error {
	return errors.New("foobar")
}
