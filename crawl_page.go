package main

import (
	"fmt"
	"net/url"
	// "strings"
)

func (config *Config) CrawlPage(rawCurrentURL string) {
	config.concurrencyControl <- struct{}{}
	defer func() {
		<-config.concurrencyControl
		config.wg.Done()
	}()

	// Get the host out of the config struct
	baseDomain := config.baseURL.Host
	currentDomain := GetDomain(rawCurrentURL)
	// fmt.Print("current domain: ", currentDomain, "\n")

	// if the base domain is not the same as the current domain, we can skip this
	if baseDomain != currentDomain {
		fmt.Printf("skipping url: %s\n", rawCurrentURL)
		return
	}

	// Adding the url to the map
	normalizedCurrentURL := NormalizeURL(rawCurrentURL)
	isFirstVisit := config.addPageVisit(normalizedCurrentURL)
	if !isFirstVisit {
		return
	} 

	// Getting the HTML page
	htmlPage, err := GetHTML(rawCurrentURL)
	if err != nil {
		_ = fmt.Errorf("error getting HTML: %s", err)
	}
	// fmt.Printf("Getting html body for url: %s\n", rawCurrentURL)

	// Getting the links in the html
	links, err := GetURLsFromHTML(htmlPage, config.baseURL.String())
	if err != nil {
		_ = fmt.Errorf("error getting links from html: %s", err)
	}
	fmt.Printf("Found %d links in the website: %s\n", len(links), rawCurrentURL)
	for _, link := range links {
		config.wg.Add(1)
		go config.CrawlPage(link)
	}

}

func GetDomain(rawURL string) string {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return ""
	}

	domain := parsedURL.Host

	return domain
}

func (cfg *Config) addPageVisit(normalizedURL string) (isFirst bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	if cfg.pages[normalizedURL] == 0 {
		cfg.pages[normalizedURL] = 1
		return true
	} else {
		cfg.pages[normalizedURL]++
		return false
	}
}
