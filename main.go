package main

import (
	"fmt"
	"net/url"
	"os"
	"sync"
)

func main() {
	argsWithoutProg := os.Args[1:]
	websiteURL := os.Args[1]
	URLstruct, err := url.Parse(websiteURL)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if len(argsWithoutProg) > 1 {
		fmt.Println("too many arguments provided, you can only provide one website link")
		os.Exit(1)
	} else if len(argsWithoutProg) == 0 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	fmt.Printf("starting crawl of: %s\n", websiteURL)

	configStruct := Config{
		pages:              make(map[string]int),
		baseURL:            URLstruct,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, 1),
		wg:                 &sync.WaitGroup{},
	}

	configStruct.wg.Add(1)
	go configStruct.CrawlPage(websiteURL)
	configStruct.wg.Wait()


}

type Config struct {
	pages              map[string]int
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}
 
