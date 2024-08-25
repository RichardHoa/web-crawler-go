package main

import (
	"net/http"
	"fmt"
	"io"
)




func GetHTML(rawURL string) (string, error) {

	resp, respErr := http.Get(rawURL)
	if respErr != nil {
		return "", respErr
	}
	if resp.StatusCode >= 400 {
		return "", fmt.Errorf("request failed with status code: %d", resp.StatusCode)
	}
	defer resp.Body.Close()
	
	body, readingHTMLErr := io.ReadAll(resp.Body)
	if readingHTMLErr != nil {
		return "", readingHTMLErr
	}
	return string(body), nil
}