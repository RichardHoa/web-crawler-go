package main

import (
	"fmt"
)

func PrintReport(pages map[string]int, baseURL string) {
	fmt.Printf("=============================\nREPORT for %s \n=============================\n", baseURL)	

	for url, count := range pages {
		fmt.Printf("Found %d internal links to %s\n", count, url)
	}

	fmt.Printf("We found a total of %d internal links\n", len(pages))
}