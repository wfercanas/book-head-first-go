package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Page struct {
	URL  string
	Size int
}

func responseSize(url string, channel chan Page) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	page := Page{URL: url, Size: len(body)}
	channel <- page
}

func main() {
	channel := make(chan Page)

	urls := []string{"https://example.com", "https://golang.org", "https://golang.org/docs"}
	for _, url := range urls {
		go responseSize(url, channel)
	}

	for range urls {
		fmt.Println(<-channel)
	}
}
