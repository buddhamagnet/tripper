package tripper

import (
	"log"

	circuit "github.com/rubyist/circuitbreaker"
)

func NewBreaker(threshold int64, breaker string, breaks map[string]func()) *circuit.Breaker {
	var cb *circuit.Breaker

	switch breaker {
	case "threshold":
		cb = circuit.NewThresholdBreaker(threshold)
	case "consecutive":
		cb = circuit.NewConsecutiveBreaker(threshold)
	default:
		log.Fatal("invalid breaker type")
	}

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
