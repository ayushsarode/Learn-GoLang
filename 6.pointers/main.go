package main

import "fmt"

func main() {
    x := 42          // Declare an integer variable
    p := &x          // Create a pointer to x

    fmt.Println("Value of x:", x)         // Prints: Value of x: 42
    fmt.Println("Address of x:", p)       // Prints: Address of x: <some memory address>
    fmt.Println("Value through p:", *p)   // Prints: Value through p: 42

    *p = 21           // Modify the value at the memory address
    fmt.Println("New value of x:", x)     // Prints: New value of x: 21
}
