package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"time"
	"log"
	"strconv"	
)
func main() {

	ClientUDP()
}

func openLogFile(path string) (*os.File, error) {
    logFile, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
    if err != nil {
        return nil, err
    }
    return logFile, nil
}

func setUpLog(){
	var test_n = ""

	if(len(os.Args) > 2){
		test_n = os.Args[2]
	}

	file, err := openLogFile("../tests/udp_logs/results_udp_"+test_n+".log")

    if err != nil { log.Fatal(err)}
	log.SetOutput(file)
    log.SetFlags(0)
}


func ClientUDP() {
	setUpLog()

	var requestTime time.Duration

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

	// Create request
	request := 2

	for i := 0; i < 10000; i++ {
		
		// prepara request & start time
		t1 := time.Now()
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

		requestTime = time.Now().Sub(t1)
		if(len(os.Args) > 1 && os.Args[1] == "teste"){
			log.Println(strconv.Itoa(int(requestTime.Nanoseconds())))
		}
	}
}
