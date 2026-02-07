package main

import (
	"fmt"
	"time"

	"game_latency_optimizer/client/forwarder"
	"game_latency_optimizer/client/health"
	"game_latency_optimizer/client/rtt"
)

func main() {
	relayAddr := "103.102.46.130:9999" // VPS relay address

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
		// 🔹 1. Direct RTT (no relay, straight to echo server)
		directRTT, err := rtt.Measure("103.102.46.130:10000")
		if err != nil {
			fmt.Println("Direct RTT: FAILED")
		} else {
			fmt.Println("Direct RTT:", directRTT)
		}

		// 🔹 2. Relay RTT (via VPS relay)
		relayRTT, err := rtt.Measure(relayAddr)
		if err != nil {
			fmt.Println("Relay UNHEALTHY")
			state.SetUnhealthy()
		} else {
			fmt.Println("Relay RTT:", relayRTT)
			fmt.Println("Relay HEALTHY")
			state.SetHealthy()
		}

		fmt.Println("----------------------------")
		time.Sleep(1 * time.Second)
	}

}
