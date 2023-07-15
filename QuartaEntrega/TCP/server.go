package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strconv"
	"tcp_module_lib/datareader"
)

func main() {
	ServerTCP()
}

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

	defer ln.Close()

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
	//Close connection
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
	}(conn)

	for {

		buffer := make([]byte, 1024)
		mLen, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("1")
			fmt.Println(err)
			os.Exit(0)
		}

		payloadAmount, err := strconv.Atoi(string(buffer[:mLen]))
		if err != nil {
			fmt.Println("2")
			fmt.Println(err)
			os.Exit(0)
		}

		r := datareader.DataReader{}.GetData(payloadAmount)

		fmt.Println(r)

		replyMsgBytes, err := json.Marshal(r)
		if err != nil {
			fmt.Println("3")
			fmt.Println(err)
			os.Exit(0)
		}

		_, err = conn.Write([]byte(replyMsgBytes))
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
	}
	// conn.Close()
}
