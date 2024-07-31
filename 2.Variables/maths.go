package main

import ("fmt"
        "math"
        "math/rand")

func maths() {
    fmt.Println("\nThis is maths.go file")
    	//math random package
	fmt.Println("My favorite number is", rand.Intn(100))
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(25))
	fmt.Println(math.Pi)
    fmt.Println(add(12,12))
}

// in this type parameter we can shortened from (x int, y int) to (x, y int)
func add(x , y int) int {
	return x + y
}