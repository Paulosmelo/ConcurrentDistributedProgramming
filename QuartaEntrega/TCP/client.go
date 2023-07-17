package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
	"log"
	//"math/rand"
)

func openLogFile(path string) (*os.File, error) {
    logFile, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
    if err != nil {
        return nil, err
    }
    return logFile, nil
}

func test_client_TCP() {	
	file, err := openLogFile("../tests/logs/results_tcp.log")
    if err != nil {
        log.Fatal(err)
    }
	log.SetOutput(file)
    log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmicroseconds)

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

	for i := 0; i < 10000; i++ {
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

		log.Println(requestTime.Nanoseconds())
	}
}

func main() {
	go test_client_TCP()

	_, _ = fmt.Scanln()
}
