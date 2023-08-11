package main

import (
	"os"
	"strconv"
	"time"
	"log"	
	"net/rpc"
	"fmt"
)

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

	file, err := openLogFile("../tests/rpc_logs/results_rpc_"+test_n+".log")

    if err != nil { log.Fatal(err)}
	log.SetOutput(file)
    log.SetFlags(0)
}

func client_RPC() {		
	setUpLog()

	var requestTime time.Duration

	// conecta ao servidor
	client, err := rpc.Dial("tcp", "localhost:1313")
	if err != nil {
		log.Fatal(err)
	}

	// fecha conexão
	defer client.Close()

	
	var reply [][]string

	// Create request
	request := 10000

	for i := 0; i < request; i++ {
		// prepara request & start time
		t1 := time.Now()

		//fmt.Println(reply)

		// invoca operação remota
		client.Call("DataReader.GetDataRPC", 2, &reply)
				
		requestTime = time.Now().Sub(t1)

		if(len(os.Args) > 1 && os.Args[1] == "teste"){
			log.Println(strconv.Itoa(int(requestTime.Nanoseconds())))
		}
	}
}

func main() {
	go client_RPC()

	_, _ = fmt.Scanln()
}
