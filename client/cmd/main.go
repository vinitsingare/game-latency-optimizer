package main

import (
	"fmt"
	"time"

	"game_latency_optimizer/client/rtt"
)

func main() {
	directAddr := "127.0.0.1:10000" // echo server
	overlayAddr := "127.0.0.1:9000" // forwarder

	for {
		directRTT, err := rtt.Measure(directAddr)
		if err != nil {
			fmt.Println("Direct error:", err)
			continue
		}

		overlayRTT, err := rtt.Measure(overlayAddr)
		if err != nil {
			fmt.Println("Overlay error:", err)
			continue
		}

		fmt.Println("Direct RTT :", directRTT)
		fmt.Println("Overlay RTT:", overlayRTT)
		fmt.Println("--------------------------------")

		time.Sleep(1 * time.Second)
	}
}
