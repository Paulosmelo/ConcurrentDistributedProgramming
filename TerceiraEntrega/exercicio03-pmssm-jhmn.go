package main

import (
	"fmt"
	"sync"
	"time"
	// "math/rand"
)

var (
	mu     sync.RWMutex
	vazio  = make(chan bool, 1)
	n      = 5
	panela = make(chan int, n)
	wg     sync.WaitGroup
)

func cliente() {

	for true {
		pegarSopa()
		time.Sleep(2 * time.Second)
	}
	defer wg.Done()
}

func cozinheiro() {

	for true {
		select {
		case <-vazio:
			encherPanela()
		default:
		}
	}

	defer wg.Done()
}

func encherPanela() {
	fmt.Println("Cozinheiro enchendo panela.")
	mu.Lock()
	for i := 0; i < n; i++ {
		panela <- 1
	}
	mu.Unlock()
}

func pegarSopa() {
	mu.RLock()
	defer mu.RUnlock()
	select {
	case <-panela:
		fmt.Println("Cliente retirou uma porção de sopa da panela.")
		comer()
	default:
		fmt.Println("Acordando cozinheiro")
		vazio <- true
	}
}

func comer() {
	fmt.Println("Cliente comeu.")
}

func main() {

	wg.Add(1)

	go cliente()
	go cozinheiro()

	wg.Wait()
}
