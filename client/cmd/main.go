package main

import (
	"fmt"
	"time"

	"game_latency_optimizer/client/config"
	"game_latency_optimizer/client/metrics"
	"game_latency_optimizer/client/probe"
	"game_latency_optimizer/client/routing"
)

func main() {
	relays := config.DefaultRelays()

	stats := make(map[string]*metrics.RTTStats)
	for i, r := range relays {
		stats[r.Name] = metrics.NewRTTStats(5)
		fmt.Printf("Initialized RTTStats for %s at index %d\n", r.Name, i)
	}

	for {
		results := probe.ProbeAll(relays, 500*time.Millisecond)

		scores := make([]routing.RouteScore, 0, len(relays))

		for i := 0; i < len(relays); i++ {
			res := <-results

			if res.Err != nil {
				fmt.Println("Probe failed for", res.RelayName, res.Err)
				continue
			}

			stats[res.RelayName].Add(res.RTT)
			avg := stats[res.RelayName].Avg()

			scores = append(scores, routing.RouteScore{
				Name: res.RelayName,
				RTT:  avg,
			})

			fmt.Println(res.RelayName, "RTT:", res.RTT, "Avg:", avg)
			time.Sleep(1 * time.Second)
		}
	}

}
