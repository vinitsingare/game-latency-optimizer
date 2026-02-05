package probe

import (
	"net"
	"time"
)

func MeasureRTT(addr string, timeout time.Duration) (time.Duration, error) {
	raddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return 0, err
	}

	conn, err := net.DialUDP("udp", nil, raddr)
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	// ⬇️ timeout protection
	conn.SetDeadline(time.Now().Add(timeout))

	start := time.Now()
	_, err = conn.Write([]byte("ping"))
	if err != nil {
		return 0, err
	}

	buf := make([]byte, 64)
	_, err = conn.Read(buf)
	if err != nil {
		return 0, err
	}

	return time.Since(start), nil
}
