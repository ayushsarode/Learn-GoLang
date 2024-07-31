package main

import "fmt"

// Define the Repeat function
func Repeat(character string) string {
    var repeated string
	for i:= 0; i< 5; i++ {
		repeated += character
	} 
    return repeated
}

// The main function, which is the entry point of the program
func main() {
    result := Repeat("a")
    fmt.Println(result) // Output: aaaaa
}
