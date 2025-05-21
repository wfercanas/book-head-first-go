package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func responseSize(url string, channel chan int) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	channel <- len(body)
}

func main() {
	channel := make(chan int)

	urls := []string{"https://example.com", "https://golang.org", "https://golang.org/docs"}
	for _, url := range urls {
		go responseSize(url, channel)
	}

	for range urls {
		fmt.Println(<-channel)
	}
}
