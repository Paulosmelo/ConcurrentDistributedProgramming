package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"tcp_module_lib/datareader"
)

func main() {

	ServerUDP()

	fmt.Scanln()
}

func ServerUDP() {
	msgFromClient := make([]byte, 1024)

	// resolve server address
	addr, err := net.ResolveUDPAddr("udp", "localhost:1313")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	// listen on udp port
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	// close conn
	defer func(conn *net.UDPConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	fmt.Println("Server UDP is ready to accept requests at port 1313...")

	for {
		// receive request
		n, addr, err := conn.ReadFromUDP(msgFromClient)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		// handle request
		HandleUDPRequest(conn, msgFromClient, n, addr)
	}
}

func HandleUDPRequest(conn *net.UDPConn, msgFromClient []byte, n int, addr *net.UDPAddr) {
	var msgToClient []byte
	var request int

	//unmarshall request
	err := json.Unmarshal(msgFromClient[:n], &request)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	// process request
	r := datareader.DataReader{}.GetData(request)

	// serialise response
	msgToClient, err = json.Marshal(r)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	// send response
	_, err = conn.WriteTo(msgToClient, addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}
