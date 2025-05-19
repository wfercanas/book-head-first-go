package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
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
	responseSize("https://example.com")
	responseSize("https://golang.org")
	responseSize("https://golang.org/docs")
}
