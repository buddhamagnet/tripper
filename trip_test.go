package tripper

import (
	"errors"
	"log"
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
	log.Fatal(failed)
}
