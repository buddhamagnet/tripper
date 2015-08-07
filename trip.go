package tripper

import (
	"github.com/rubyist/circuitbreaker"
)

func tripThreshold(threshold int64, trip, fail, reset, ready, event func()) *circuit.Breaker {
	cb := circuit.NewThresholdBreaker(threshold)

	events := cb.Subscribe()
	go func() {
		for {
			e := <-events
			switch e {
			case circuit.BreakerTripped:
				trip()
			case circuit.BreakerReset:
				reset()
			case circuit.BreakerFail:
				fail()
			case circuit.BreakerReady:
				ready()
			default:
				event()
			}
		}
	}()
	return cb
}
