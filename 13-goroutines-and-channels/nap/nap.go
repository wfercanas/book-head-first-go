package main

import (
	"fmt"
	"time"
)

func nap(name string, delay int) {
	for i := 0; i < delay; i++ {
		fmt.Println(name, "is sleeping")
		time.Sleep(time.Second)
	}
	fmt.Println(name, "has waken up!")
}

func sender(channel chan string) {
	nap("sender", 2)
	fmt.Println("*** Ready to send the first value ***")
	channel <- "a"
	fmt.Println("*** first value sent ***")
	fmt.Println("*** Ready to send the second value ***")
	channel <- "b"
	fmt.Println("*** second value sent ***")
}

func main() {
	channel := make(chan string)
	go sender(channel)
	nap("main", 5)
	fmt.Println(<-channel)
	fmt.Println(<-channel)
}
