package probe

import (
	"time"

	"game_latency_optimizer/client/config"
)

func ProbeAll(
	relays []config.Relay,
	timeout time.Duration,
) <-chan Result {

	results := make(chan Result)

	for _, relay := range relays {
		go func(r config.Relay) {
			rtt, err := MeasureRTT(r.Addr, timeout)
			results <- Result{
				RelayName: r.Name,
				RTT:       rtt,
				Err:       err,
			}
		}(relay)
	}

	return results
}
