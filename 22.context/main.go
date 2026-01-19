package main

import (
	"context"
	"fmt"
)

// "context"
// "fmt"
// "time"

// func dosomething(ctx context.Context) {
// 	select {
// 	case <- time.After(5 * time.Second):
// 		fmt.Println("Finished work")
// 	case <- ctx.Done():
// 		fmt.Println("Canceled:", ctx.Err())
// 	}
// }

func main() {
	// ctx, cancel := context.WithTimeout(context.Background(), 6 * time.Second)
	// defer cancel()


	// dosomething(ctx)

    ctx := context.Background() // root context
    fmt.Println("Context:", ctx)

}


