package main

import (
	"fmt"
	"time"

	"game_latency_optimizer/client/forwarder"
	"game_latency_optimizer/client/health"
	"game_latency_optimizer/client/rtt"
)

func main() {
	relayAddr := "127.0.0.1:9999" // VPS relay address

	state := health.NewState()

	// Start forwarder
	f := forwarder.New(":7000", relayAddr, state)
	go func() {
		if err := f.Start(); err != nil {
			panic(err)
		}
	}()

	// Health loop (control plane)
	for {
		_, err := rtt.Measure(relayAddr)
		if err != nil {
			fmt.Println("Relay UNHEALTHY")
			state.SetUnhealthy()
		} else {
			fmt.Println("Relay HEALTHY")
			state.SetHealthy()
		}
		time.Sleep(1 * time.Second)
	}
}
