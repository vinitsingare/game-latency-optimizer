package forwarder

import (
	"log"
	"net"
	"sync"

	"game_latency_optimizer/client/health"
)

type Forwarder struct {
	localAddr string // where local apps send packets
	relayAddr string // VPS relay address
	conn      *net.UDPConn
	relayConn *net.UDPConn
	clientMap map[string]*net.UDPAddr
	mu        sync.Mutex
	health    *health.State
}

func New(localAddr, relayAddr string, h *health.State) *Forwarder {
	return &Forwarder{
		localAddr: localAddr,
		relayAddr: relayAddr,
		health:    h,
		clientMap: make(map[string]*net.UDPAddr),
	}
}

func (f *Forwarder) Start() error {
	laddr, err := net.ResolveUDPAddr("udp", f.localAddr)
	if err != nil {
		return err
	}

	f.conn, err = net.ListenUDP("udp", laddr)
	if err != nil {
		return err
	}

	raddr, err := net.ResolveUDPAddr("udp", f.relayAddr)
	if err != nil {
		return err
	}

	f.relayConn, err = net.DialUDP("udp", nil, raddr)
	if err != nil {
		return err
	}

	log.Println("Forwarder listening on", f.localAddr)
	log.Println("Forwarding to relay", f.relayAddr)

	go f.readFromClients()
	go f.readFromRelay()

	return nil
}

func (f *Forwarder) readFromClients() {
	buf := make([]byte, 4096)
	for {
		n, clientAddr, err := f.conn.ReadFromUDP(buf)
		if err != nil {
			continue
		}

		// 🔴 CONTROL → DATA PLANE GATE
		if !f.health.IsHealthy() {
			// drop packet when relay unhealthy
			continue
		}

		f.mu.Lock()
		f.clientMap[clientAddr.String()] = clientAddr
		f.mu.Unlock()

		_, _ = f.relayConn.Write(buf[:n])
	}
}

func (f *Forwarder) readFromRelay() {
	buf := make([]byte, 4096)

	for {
		n, _, err := f.relayConn.ReadFromUDP(buf)
		if err != nil {
			continue
		}

		// send response back to all known clients (simple version)
		f.mu.Lock()
		for _, clientAddr := range f.clientMap {
			f.conn.WriteToUDP(buf[:n], clientAddr)
		}
		f.mu.Unlock()
		log.Println("Forwarder: packet from relay")

	}
}
