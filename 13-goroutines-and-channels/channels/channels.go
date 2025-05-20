package main

import "fmt"

func greeting(channel chan string) {
	channel <- "hi"
}

func main() {
	channel := make(chan string)
	go greeting(channel)
	greet := <-channel
	fmt.Println(greet)
}
