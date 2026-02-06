package main

import (
	"log"
	"net"
)

func main() {
	listenAddr := ":9000" // VPS listens here

	conn, err := net.ListenPacket("udp", listenAddr)
	if err != nil {
		log.Fatal("Failed to listen:", err)
	}
	defer conn.Close()

	log.Println("UDP relay listening on", listenAddr)

	buf := make([]byte, 4096)

	for {
		n, clientAddr, err := conn.ReadFrom(buf)
		if err != nil {
			continue
		}

		// 🔴 TEMP destination (we’ll make this dynamic later)
		targetAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:10000")
		if err != nil {
			continue
		}

		log.Println("Relay: sending response back to client")
		go forwardPacket(conn, clientAddr, targetAddr, buf[:n])

	}
}

func forwardPacket(
	conn net.PacketConn,
	clientAddr net.Addr,
	targetAddr *net.UDPAddr,
	data []byte,
) {
	targetConn, err := net.DialUDP("udp", nil, targetAddr)
	if err != nil {
		return
	}
	defer targetConn.Close()

	// Send to target
	_, err = targetConn.Write(data)
	if err != nil {
		return
	}

	resp := make([]byte, 4096)
	n, _, err := targetConn.ReadFrom(resp)
	if err != nil {
		return
	}

	// Send response back to client
	conn.WriteTo(resp[:n], clientAddr)
	log.Println("Relay: packet from client", clientAddr)

}
