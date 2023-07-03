package main

import (
	"fmt"
	"time"
	"sync"
	"math/rand"
)

var (
	ocupado = make(chan bool, 1)
	fila = make(chan int, 5)
	wg sync.WaitGroup
)

func cliente(){
	var i = 1
	rand.Seed(time.Now().UnixNano())
    
	for true{
    	n := rand.Intn(8)
		time.Sleep(time.Duration(n)* time.Second)
		DesejaCortarCabelo(i)
		i++
	}
	defer wg.Done()
}

func barbeiro(){

	for true {
		select{
		case ocupado <- true:
			CortarCabelo()
		default:

		}
	}
	defer wg.Done()
}

func CortarCabelo(){
	select{
	case i:=<-fila:
		fmt.Println("cliente ", i, " está cortando o cabelo.")
		time.Sleep(5*time.Second)
		<-ocupado
	default:
		fmt.Println("Barbeiro cochilando...")
		<-ocupado
	}
}

func DesejaCortarCabelo(i int){
	select{
	case fila<-i:
		fmt.Println(i, " deseja cortar cabelo")
	default:
			fmt.Println("Cliente ", i, " saiu.")
	}
}

func main(){
	
	wg.Add(1)

	go 	cliente()
	go barbeiro()

	wg.Wait()
}