package main

import (
	"fmt"
	"net"
	"os"
	"tcp_module_lib/datareader"
	"net/rpc"
)

func server() {
	datareader := new(datareader.DataReaderRPC)

	// Cria server RPC
	server := rpc.NewServer()
	err := server.RegisterName("DataReader", datareader)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	// Cria listener
	ln, err := net.Listen("tcp", "localhost:1313")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	defer func(ln net.Listener) {
		var err = ln.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(ln)

	// aguarda por requests
	fmt.Println("Server is ready ...")
	server.Accept(ln)
}

func main() {

	go server()

	_, _ = fmt.Scanln()
}