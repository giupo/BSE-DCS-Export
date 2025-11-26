package main

import (
	"flag"
	"fmt"
)

func main() {
	udpPort := flag.String("udp", "6666", "Porta UDP da ascoltare")
	tcpPort := flag.String("tcp", "6666", "Porta TCP da ascoltare")
	apiPort := flag.Int("http", 8080, "Porta TCP da ascoltare")
	flag.Parse()

	var store *DataStore = NewDataStore()

	if *udpPort == "" && *tcpPort == "" {
		fmt.Println("Usa: app -udp 6666  oppure:  app -tcp 6666")
		return
	}

	if *udpPort != "" {
		u := UDPServer{Port: *udpPort, Store: store}
		u.Start()
	}

	if *tcpPort != "" {
		t := TCPServer{Port: *tcpPort, Store: store}
		t.Start()
	}

	StartAPI(store, *apiPort)
}
