package main

import (
	"net/url"
	"strings"
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