package main

import (
	"fmt"
	"time"
)

// Go provides channels that you can use for bidirectional communication between goroutines.

func main() {
	msg := make(chan string)

	go greet(msg)

	time.Sleep(2 * time.Second)

	greeting := <- msg

	time.Sleep(2 * time.Second)

	fmt.Println("Greeting received")

	fmt.Println(greeting)

    _, ok := <-msg
    if ok {
        fmt.Println("Channel is open!")
    } else {
        fmt.Println("Channel is closed!")
    }

	
}

func greet(ch chan string) {
	fmt.Println("Greeter waiting to send greeting!")

	ch <- "Hello ayush"

	close(ch)

	fmt.Println("Greeter Completed")
}