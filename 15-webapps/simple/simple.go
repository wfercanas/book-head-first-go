package main

import (
	"log"
	"net/http"
)

func write(writer http.ResponseWriter, message []byte) {
	_, err := writer.Write(message)
	if err != nil {
		log.Fatal(err)
	}
}

func englishHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, []byte("Hello"))
}

func hindiHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, []byte("Namaste"))
}

func frenchHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, []byte("Salut"))
}

func main() {
	http.HandleFunc("/hello", englishHandler)
	http.HandleFunc("/hindi", hindiHandler)
	http.HandleFunc("/salut", frenchHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
