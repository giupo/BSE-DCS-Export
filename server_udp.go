package main

import (
	"fmt"
	"net"
)

// ---------------- UDPServer ----------------
type UDPServer struct {
	Port       string
	BufferSize uint
	Store      *DataStore
}

func (s *UDPServer) Start() error {
	addr := ":" + s.Port
	s.Store = NewDataStore()
	conn, err := net.ListenPacket("udp", addr)
	s.BufferSize = 65000
	if err != nil {
		return fmt.Errorf("errore apertura socket UDP: %w", err)
	}
	fmt.Println("UDP server in ascolto su", addr)

	buf := make([]byte, s.BufferSize)
	go func() {
		defer conn.Close()
		for {
			n, _, err := conn.ReadFrom(buf)
			if err != nil {
				fmt.Println("Errore ricezione UDP:", err)
				continue
			}

			updateData(s.Store, buf[:n])
		}
	}()
	return nil
}
