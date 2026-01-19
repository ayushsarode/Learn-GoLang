package main

import (
	"fmt"
	"time"
)

func main() {
	ticket := time.NewTicker(1 * time.Second)
	defer ticket.Stop()
	done := make(chan bool)

	go func() {
		for i := 10; i >= 0; i-- {
			select {
			case <- ticket.C:
				fmt.Println(i)
			case <- done:
				return
			}
		}
		done <- true
	}()

	fmt.Println("start")
	<-done
	fmt.Println("finish")

}
