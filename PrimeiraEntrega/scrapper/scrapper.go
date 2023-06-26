package main

import (
    "github.com/gocolly/colly"
    "log"
    "os"
)


func get_fighters_links( collector colly.Collector){
    
    for r := 'a'; r <'z'; r++ {        
		get_links_by_letter(string(r), collector)       
	}


}

func get_links_by_letter(letter string, collector colly.Collector){
    //var wg sync.WaitGroup

    // Find and print all links
    collector.OnHTML("td.b-statistics__table-col", func(e *colly.HTMLElement) {
        links := e.ChildAttrs("a", "href")

        if(len(links)> 0){
            val := string(links[0])
            
            append_to_file(val)
        
        }
    })
    collector.Visit("http://www.ufcstats.com/statistics/fighters?char="+letter)
}

func append_to_file(val string){
    f, err := os.OpenFile("data.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
    if err != nil {
        panic(err)
    }
    
    defer f.Close()
    
    if _, err = f.WriteString(val + "\n"); err != nil {
        panic(err)
    }
    if err != nil {
        log.Fatal(err)
    }
}

func main() {

    f, err := os.Create("data.txt")

    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    collector := colly.NewCollector(
        colly.AllowedDomains("www.ufcstats.com"),
    )	
   
    get_fighters_links(*collector)

    // var i string
    // fmt.Scan(&i, "end")
}	