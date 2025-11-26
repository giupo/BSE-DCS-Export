package main

import (
	"fmt"
	"net"
)

type TCPServer struct {
	Port       string
	BufferSize uint
	Store      *DataStore
	Buffer     []byte
}

func (s *TCPServer) Start() error {
	addr := ":" + s.Port
	s.BufferSize = 1024 * 1024
	s.Buffer = make([]byte, s.BufferSize)
	s.Store = NewDataStore()
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("errore apertura socket TCP: %w", err)
	}

	fmt.Println("TCP server in ascolto su", addr)
	go s.acceptLoop(ln)
	return nil
}

func (s *TCPServer) acceptLoop(ln net.Listener) {
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Errore Accept TCP:", err)
			continue
		}

		go s.handleConn(conn)
	}
}

func (s *TCPServer) handleConn(conn net.Conn) {
	defer conn.Close()

	for {
		n, err := conn.Read(s.Buffer)
		if err != nil {
			fmt.Println("Connessione TCP chiusa:", err)
			return
		}
		updateData(s.Store, s.Buffer[:n])
	}
}
