package rtt

import (
	"net"
	"time"
)

func Measure(addr string) (time.Duration, error) {
	conn, err := net.Dial("udp", addr)
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	// 🔴 THIS LINE IS CRITICAL
	conn.SetReadDeadline(time.Now().Add(2 * time.Second))

	start := time.Now()

	_, err = conn.Write([]byte("ping"))
	if err != nil {
		return 0, err
	}

	buf := make([]byte, 1024)
	_, err = conn.Read(buf)
	if err != nil {
		return 0, err
	}

	return time.Since(start), nil
}
