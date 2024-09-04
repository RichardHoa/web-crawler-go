package main

import (
	"fmt"
	"sort"
)

func PrintReport(pages map[string]int, baseURL string) {
	fmt.Printf("=============================\nREPORT for %s \n=============================\n", baseURL)

	// Convert the map to a slice of key-value pairs
	type kv struct {
		Key   string
		Value int
	}

	var sortedPages []kv
	for url, count := range pages {
		sortedPages = append(sortedPages, kv{url, count})
	}

	// Sort the slice by count in descending order
	sort.Slice(sortedPages, func(i, j int) bool {
		return sortedPages[i].Value > sortedPages[j].Value
	})

	// Print the sorted pages
	for _, kv := range sortedPages {
		fmt.Printf("Found %d internal links to %s\n", kv.Value, kv.Key)
	}

	fmt.Printf("We found a total of %d internal links\n", len(pages))
}
