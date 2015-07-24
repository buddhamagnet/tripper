package tripper

import (
	"github.com/rubyist/circuitbreaker"
)

func tripThreshold(threshold int64, breaks map[string]func()) *circuit.Breaker {
	cb := circuit.NewThresholdBreaker(threshold)

	events := cb.Subscribe()
	go func() {
		for {
			e := <-events
			switch e {
			case circuit.BreakerTripped:
				breakMe("trip", breaks)
			case circuit.BreakerReset:
				breakMe("reset", breaks)
			case circuit.BreakerFail:
				breakMe("fail", breaks)
			case circuit.BreakerReady:
				breakMe("ready", breaks)
			default:
				breakMe("event", breaks)
			}
		}
	}()
	return cb
}

func breakMe(event string, breaks map[string]func()) {
	if callback, found := breaks[event]; found {
		callback()
	}
}
