package main

import (
    "fmt"	
    "sync"
    "github.com/gocolly/colly"
    "bufio"
)

var fighters_page []string



func get_fighters_links(c chan int, collector colly.Collector){
    var wg sync.WaitGroup

    for r := 'a'; r <'z'; r++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
		    get_details_page_by_letter(string(r), c, collector)
        }()
	}


    wg.Wait()
}

func get_details_page_by_letter(letter string, c chan int, collector colly.Collector){
    // Find and print all links
    collector.OnHTML("td.b-statistics__table-col", func(e *colly.HTMLElement) {
        links := e.ChildAttrs("a", "href")
        name := e.ChildText("a")
        fmt.Println(name)
        if(len(links)> 0){
            fighters_page =  append(fighters_page, links[0])            
        }else{
          
        }
    })
    collector.Visit("http://www.ufcstats.com/statistics/fighters?char="+letter)
}

func main() {
    c := make(chan int, 1)

    collector := colly.NewCollector(
        colly.AllowedDomains("www.ufcstats.com"),
    )	
    var wg sync.WaitGroup
    wg.Add(1)
    go func() {
        defer wg.Done()
        go get_fighters_links(c, *collector)
    }()

    wg.Wait()

    fmt.Println("teste")
    for _, v := range fighters_page {
        fmt.Println(v)
    }

    var i string
    fmt.Scan(&i)

}	