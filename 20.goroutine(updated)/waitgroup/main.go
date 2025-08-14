package main

import (
	"fmt"
	"sync"
)

//Blocking the method for 2 seconds might not create any problems. But at the production level, where each millisecond is vital, millions of concurrent requests can create a huge problem.

// You can solve this problem using sync.WaitGroup

//This ensures both goroutines finish before main exits, but their order is still not guaranteed.
func main(){
	var wg sync.WaitGroup
	ch := make(chan bool)

	wg.Add(2)

	go func() {
		defer wg.Done()
		fmt.Println("Hello world!")
		ch <- true
	}()

	go func() {
		defer wg.Done()
		<- ch
		fmt.Println("Goodbye!")
	}()

	wg.Wait()
}