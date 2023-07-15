package datareader

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type DataReaderTCP struct{}

func (DataReaderTCP) GetData(quantLinhas int) [][]string {
	var linhas = make([][]string, quantLinhas)
	fmt.Println()
	f, err := os.Open("./datareader/data.csv")
	if err != nil {
		log.Fatal(err)
	}

	csvReader := csv.NewReader(f)

	for i := 0; i <= quantLinhas; i++ {
		// read csv values using csv.Reader
		data, err := csvReader.Read()
		if err != nil {
			log.Fatal(err)
		}
		if i > 0 {
			linhas[i-1] = append(linhas[i-1], data...)
		}
	}
	f.Close()

	return linhas[:]
}

// func main() {
// 	linhas := getData(5)

// 	fmt.Println(linhas)

// }
