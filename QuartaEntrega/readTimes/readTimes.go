package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	var times = make([]float64, 10000)
	var row = make([]string, 3)
	var somatorio float64 = 0
	var media float64 = 0
	f, err := os.Open("../tests/logs/results_udp.log")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	i := 0
	for scanner.Scan() {
		row = strings.Split(scanner.Text(), " ")
		if row[3] == "teste" {
			times[i], err = strconv.ParseFloat(row[4], 64)
			somatorio += times[i]
			fmt.Println(somatorio)
			i++
		}
	}
	//calculo da media
	media = somatorio / 10000
	fmt.Println("Média: ", media)
	var somatorioDP float64 = 0
	var dp float64 = 0

	//calculo do desvio padrão
	for i := 0; i < 10000; i++ {
		somatorioDP += math.Pow(math.Abs(times[i]-media), 2)
	}
	dp = math.Sqrt(somatorio / 10000)

	fmt.Println("Desvio padrão:", dp)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
