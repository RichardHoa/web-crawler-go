package main

import (
	"fmt"
	"net/url"
	"os"
	"sync"
	"strconv"
)

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) > 4 {
		fmt.Println("too many arguments provided, expected command: go run . <websiteURL> <maxCocurrencies> <maxPages>")
		os.Exit(1)
	} else if len(argsWithoutProg) == 0 {
		fmt.Println("no website provided, expected command: go run . <websiteURL> <maxCocurrencies> <maxPages>")
		os.Exit(1)
	}
	
	websiteURL := os.Args[1]
	maxCocurrencies, err:= strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("error converting string to int: %s", err)
		os.Exit(1)
	}
	maxPages, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Printf("error converting string to int: %s", err)
		os.Exit(1)
}
	fmt.Printf("Website url %s max cocurrencies %d max pages %d\n", websiteURL,maxCocurrencies, maxPages)
	URLstruct, err := url.Parse(websiteURL)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("starting crawl of: %s\n", websiteURL)

	configStruct := Config{
		pages:              make(map[string]int),
		baseURL:            URLstruct,
		maxPages:           maxPages,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxCocurrencies),
		wg:                 &sync.WaitGroup{},
	}

	configStruct.wg.Add(1)
	go configStruct.CrawlPage(websiteURL)
	configStruct.wg.Wait()

	PrintReport(configStruct.pages, websiteURL)

}

type Config struct {
	pages              map[string]int
	baseURL            *url.URL
	maxPages           int
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}
