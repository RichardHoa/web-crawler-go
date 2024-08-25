package main

import (
	"net/url"
	"strings"
	"golang.org/x/net/html"
	// "fmt"
)

func NormalizeURL(inputURL string) string {
	urlStruct, err := url.Parse(inputURL)
	if err != nil {
		// Handle the error
		return ""
	}
	normalizedPath := strings.TrimSuffix(urlStruct.Path, "/")

	return urlStruct.Host + normalizedPath
}


func GetURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	doc, err := html.Parse(strings.NewReader(htmlBody))
	
	if err != nil {
		return []string{}, err
	}
	var hrefList []string

	// Recursive function to traverse the HTML tree
	var traverseHTML func(*html.Node)

	traverseHTML = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "a" {
			for _, attribute := range node.Attr {
				if attribute.Key == "href" {
					aHrefAttribute, err := url.Parse(attribute.Val)
					if err == nil && !aHrefAttribute.IsAbs() {
						baseURL, err := url.Parse(rawBaseURL)
						if err == nil {
							attribute.Val = baseURL.ResolveReference(aHrefAttribute).String()
						}
					}
					hrefList = append(hrefList, attribute.Val)
					break
				}
			}
		}
		// Recursively traverse child nodes
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			traverseHTML(child)
		}
	}

	// Start traversing from the root node
	traverseHTML(doc)

	return hrefList, nil
}
