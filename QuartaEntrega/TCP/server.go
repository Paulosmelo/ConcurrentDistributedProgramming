package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"tcp_module_lib/datareader"
)

func ServerTCP() {
	r, err := net.ResolveTCPAddr("tcp", "localhost:1313")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	ln, err := net.ListenTCP("tcp", r)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	fmt.Println("Server listening on:", r)

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Printf(err.Error())
			os.Exit(0)
		}
		go HandleTCPConnection(conn)
	}
}

func HandleTCPConnection(conn net.Conn) {
	var msgFromClient  int

	//Close connection
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
	}(conn)

	// Create coder/decoder
	jsonDecoder := json.NewDecoder(conn)
	jsonEncoder := json.NewEncoder(conn)


	for{
	
		err := jsonDecoder.Decode(&msgFromClient)
		if err != nil && err.Error() == "EOF" {
			//conn.Close()
			// no further requests
			break
		}

		if err != nil {
			fmt.Println(err)
			os.Exit(0) 
		}

		// Process request
		r := datareader.DataReader{}.GetData(msgFromClient)

		// Create response
		msgToClient := r

		// Serialise and send response to client
		err = jsonEncoder.Encode(msgToClient)
		if err != nil {
			os.Exit(0)
			break
		}
	}
	// conn.Close()
}

func main() {
	ServerTCP()

	fmt.Scanln()
}