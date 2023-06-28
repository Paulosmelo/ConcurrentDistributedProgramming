package main

import (
    "github.com/gocolly/colly"
    "os"
   "fmt"
    "sync"
)

var (
    fighters_links = make(chan string, 50000)
    c = make(chan int, 1)
    wg sync.WaitGroup
)

func get_fighters_links( collector colly.Collector){
    
    for r := 'a'; r <'z'; r++ {        
		get_links_by_letter(string(r), collector)       
	}
}

func get_links_by_letter(letter string, collector colly.Collector){
    collector.OnHTML("td.b-statistics__table-col", func(e *colly.HTMLElement) {
        links := e.ChildAttrs("a", "href")

        if(len(links)> 0){
            val := string(links[0])
            fighters_links <- val
        }
    })
    collector.Visit("http://www.ufcstats.com/statistics/fighters?char="+letter)
}

func append_to_file(batchSize int){

    f, err := os.OpenFile("data.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
    if err != nil {
        panic(err)
    }
    
    defer f.Close()
    
    c <- 1

    for i := 0; i < batchSize; i++ {
        t := <- fighters_links
        fmt.Println(i)
        f.WriteString(t + "\n");
    }
        
    <-c
}

func main() {

    os.Remove("data.txt")

    f, err := os.Create("data.txt")

    if err != nil{}

    defer f.Close()

    collector := colly.NewCollector(
        colly.AllowedDomains("www.ufcstats.com"),
    )	
   
    get_fighters_links(*collector)

    threads := 8
    max := len(fighters_links)

    //fmt.Println(max)
   
    batchSize := max/threads
    rest := max%threads
    fmt.Println(rest)

    for i := 0; i < threads; i++ {
        wg.Add(1)
        
        if(i == threads-1){
            batchSize = batchSize + rest
        }

        go func(){
            defer wg.Done()
            append_to_file(batchSize)
        }()
    }

    wg.Wait()
}	