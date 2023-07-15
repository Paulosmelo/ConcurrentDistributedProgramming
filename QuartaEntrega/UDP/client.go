package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
)

func main() {

	ClientUDP(10)
}

func ClientUDP(n int) {
	var response [][]string
	// resolve server address
	addr, err := net.ResolveUDPAddr("udp", "localhost:1313")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	// connect to server -- does not create a connection
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	// create coder/decoder
	decoder := json.NewDecoder(conn)
	encoder := json.NewEncoder(conn)

	// Close connection
	defer func(conn *net.UDPConn) {
		err := conn.Close()
		if err != nil {
		}
	}(conn)

	for i := 0; i < n; i++ {
		// Create request
		request := 1

		// Serialise and send request
		err = encoder.Encode(request)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		// Receive response from servidor
		err = decoder.Decode(&response)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		fmt.Println(response)
	}
}
