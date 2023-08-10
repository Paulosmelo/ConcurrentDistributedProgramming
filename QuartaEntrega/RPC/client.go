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

func client_TCP() {	
	file, err := openLogFile("../tests/logs/results_rpc.log")
    if err != nil {
        log.Fatal(err)
    }
	log.SetOutput(file)
    log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmicroseconds)

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
			log.Println(os.Args[1]+ strconv.Itoa(int(requestTime.Nanoseconds())))
		}else{
			log.Println(strconv.Itoa(int(requestTime.Nanoseconds())))
		}
	}
}

func main() {
	go client_TCP()

	_, _ = fmt.Scanln()
}
