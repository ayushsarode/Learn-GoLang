package main

import (
	"fmt"
	"strings"
)

func main() {
	myString := "Good morning";
	var indexed = myString[0]
	fmt.Print(indexed)
	for i, v := range myString{
		fmt.Println(i, v)
	}
	fmt.Print("\n The lenght of 'myString' is %v\n\n", len(myString))


	str := "Hello, 世界" // Mixed ASCII and Unicode characters

	// Bytes: underlying data of a string
	bytes := []byte(str)
	fmt.Println("String:", str)
	fmt.Println("Bytes:", bytes) // Each character's byte representation
	fmt.Println("Bytes as characters:")
	for _, b := range bytes {
		fmt.Printf("%c ", b) // Print bytes as characters (may show garbage for multi-byte characters)
	}
	fmt.Println()

	// Runes: Unicode code points
	runes := []rune(str)
	fmt.Println("Runes:", runes) // Unicode code points for each character
	fmt.Println("Runes as characters:")
	for _, r := range runes {
		fmt.Printf("%c ", r) // Print the actual characters
	}
	fmt.Println()

	// Length of string
	fmt.Println("Length of string (bytes):", len(str)) // Counts bytes
	fmt.Println("Number of runes (characters):", len(runes)) // Counts Unicode characters



	var strSlice = []string{"s", "u", "b", "s", "c"}
	var strBuilder strings.Builder

	for i := range strSlice{
		strBuilder.WriteString(strSlice[i])
	}

	var catStr = strBuilder.String()
	fmt.Printf("\n%v", catStr)
}