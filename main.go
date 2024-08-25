package main

import (
	"fmt"
	"os"	
)

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) > 1 {
		fmt.Println("too many arguments provided, you can only provide one website link")	
		os.Exit(1)	
	} else if len(argsWithoutProg) == 0 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	fmt.Printf("starting crawl of: %s\n", os.Args[1])

	htmlBody, getHTMLErr := GetHTML(os.Args[1])
	if getHTMLErr != nil {
		fmt.Printf("error getting HTML: %s\n", getHTMLErr)
		os.Exit(1)
	}
	fmt.Printf("HTML body:\n%s\n", htmlBody)
}