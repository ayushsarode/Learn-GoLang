package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// func main() {
// 	var wg sync.WaitGroup

// 	wg.Add(1)
// 	go func(){
// 		defer wg.Done()
// 		for i:=0; i < 50; i++ {
// 			fmt.Println("from first func",i)
// 			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

// 		}
// 	}()

// 	wg.Add(1)
// 	go func ()  {
// 		defer wg.Done()
// 		for i:= 0; i < 50; i++{
// 			fmt.Println("from second func" ,i)
// 			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

// 		}
// 	}()


// 	wg.Wait()

// 	fmt.Println("Done")
// }


func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i ++ {
		wg.Add(1)

		go func(id int){
			defer wg.Done()
			

			for j := 0; j < 20; j++ {
				fmt.Printf("Worker: %d, Task %d \n", id, j)
				time.Sleep(time.Duration(rand.Intn(50)) * time.Millisecond)
			}
			fmt.Printf("Worker %d finished \n", id )
		}(i)
	}

	wg.Wait()
	fmt.Print("All workers done")
}