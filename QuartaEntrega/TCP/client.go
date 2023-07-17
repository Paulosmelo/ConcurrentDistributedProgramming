package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
	"log"	
)

func openLogFile(path string) (*os.File, error) {
    logFile, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
    if err != nil {
        return nil, err
    }
    return logFile, nil
}

func client_TCP() {	
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


	var response [][]string

	// create coder/decoder
	decoder := json.NewDecoder(conn)
	encoder := json.NewEncoder(conn)

	// Create request
	request := 2

	for i := 0; i < 10000; i++ {
		// prepara request & start time
		t1 := time.Now()

		// Serialise and send request
		err = encoder.Encode(&request)
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


		requestTime = time.Now().Sub(t1)
		var arg = ""
		if(len(os.Args) > 1){
			arg = os.Args[1]
		}

		log.Println(arg + strconv.Itoa(int(requestTime.Nanoseconds())))
	}
}

func main() {
	go client_TCP()

	_, _ = fmt.Scanln()
}
