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
	for _, r := range relays {
		stats[r.Name] = metrics.NewRTTStats(5)
	}

	for {
		var scores []routing.RouteScore

		for _, r := range relays {
			rtt, err := probe.MeasureRTT(r.Addr)
			if err != nil {
				fmt.Println("Error probing", r.Name, err)
				continue
			}

			stats[r.Name].Add(rtt)
			avg := stats[r.Name].Avg()

			scores = append(scores, routing.RouteScore{
				Name: r.Name,
				RTT:  avg,
			})

			fmt.Println(r.Name, "RTT:", rtt, "Avg:", avg)
		}

		best := routing.ChooseBest(scores)
		fmt.Println("👉 Best Route:", best.Name, best.RTT)
		fmt.Println("--------------------------------")

		time.Sleep(1 * time.Second)
	}
}
