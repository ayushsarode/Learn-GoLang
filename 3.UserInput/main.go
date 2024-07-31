package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our restro: ")

	// comma ok || err ok syntax

	input, _ := reader.ReadString('\n')
	fmt.Println("Rating of this restro is ", input)
	fmt.Printf("Type of this variable is %T ", input)
}