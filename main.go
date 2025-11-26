package main

import (
	"flag"
	"fmt"
)

func main() {
	udpPort := flag.String("udp", "", "Porta UDP da ascoltare")
	tcpPort := flag.String("tcp", "", "Porta TCP da ascoltare")
	flag.Parse()

	if *udpPort == "" && *tcpPort == "" {
		fmt.Println("Usa: app -udp 6666  oppure:  app -tcp 6666")
		return
	}

	if *udpPort != "" {
		u := UDPServer{Port: *udpPort}
		u.Start()
		StartAPI(u.Store, 8080)
	}

	if *tcpPort != "" {
		t := TCPServer{Port: *tcpPort}
		t.Start()
		StartAPI(t.Store, 8080)
	}

	// Evita che l’app finisca

}
