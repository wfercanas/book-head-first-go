package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func responseSize(url string) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(url, len(body))

}

func main() {
	go responseSize("https://example.com")
	go responseSize("https://golang.org")
	go responseSize("https://golang.org/docs")
	time.Sleep(time.Second)
}
