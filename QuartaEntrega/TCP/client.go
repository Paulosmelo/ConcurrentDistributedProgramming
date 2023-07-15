package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
	//"math/rand"
)

func clientTCP() {
	var requestTime time.Duration

	r, err := net.ResolveTCPAddr("tcp", "localhost:1313")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	conn, err := net.DialTCP("tcp", nil, r)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	// fecha conexão
	defer func(conn *net.TCPConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	// prepara request & start time
	t1 := time.Now()

	_, err = conn.Write([]byte(strconv.Itoa(2)))
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	buffer := make([]byte, 1024)
	mLen, err := conn.Read(buffer)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	var feedback = make([][]string, 1)
	json.Unmarshal(buffer[:mLen], &feedback)
	fmt.Println(feedback)

	requestTime = time.Now().Sub(t1)

	fmt.Printf("Total Duration: %v", requestTime)
}

func main() {
	go clientTCP()

	_, _ = fmt.Scanln()
}
