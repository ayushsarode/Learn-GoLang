package main

import "fmt"

func Hello() string {
	return ("Hello World from go")
}

func main() {
	fmt.Println(Hello())
}