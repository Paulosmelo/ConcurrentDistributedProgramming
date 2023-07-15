// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net"
// 	"os"
// 	"strconv"
// 	"time"
// 	//"math/rand"
// )

// func clientTCP() {
// 	times := [1000]time.Duration{}
// 	var SAMPLE_SIZE = 1000

// 	r, err := net.ResolveTCPAddr("tcp", "localhost:1313")
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(0)
// 	}

// 	conn, err := net.DialTCP("tcp", nil, r)
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(0)
// 	}

// 	// fecha conexão
// 	defer func(conn *net.TCPConn) {
// 		err := conn.Close()
// 		if err != nil {

// 		}
// 	}(conn)

// 	// loop
// 	for i := 0; i < SAMPLE_SIZE; i++ {

// 		// rand.Seed(time.Now().UnixNano())
// 		// var random = rand.Intn(20)

// 		// prepara request & start time
// 		t1 := time.Now()

// 		_, err = conn.Write([]byte(strconv.Itoa(1)))
// 		if err != nil {
// 			fmt.Println(err)
// 			os.Exit(0)
// 		}

// 		buffer := make([]byte, 1024)
// 		mLen, err := conn.Read(buffer)
// 		if err != nil {
// 			fmt.Println(err)
// 			os.Exit(0)
// 		}

// 		var feedback [20]int
// 		json.Unmarshal(buffer[:mLen], &feedback)
// 		fmt.Println(feedback)

// 		times[i] = time.Now().Sub(t1)
// 	}

// 	totalTime := time.Duration(0)
// 	for i := range times {
// 		totalTime += times[i]
// 	}
// 	fmt.Printf("Total Duration: %v [%v]", totalTime, SAMPLE_SIZE)
// }

// func main() {
// 	go clientTCP()

// 	_, _ = fmt.Scanln()
// }
