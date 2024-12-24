package main

import (
	"errors"
	"fmt"
)

func main() {

	var numerator int = 12;
	var denominator int = 2;
	var result, remainder, err = intDivision(numerator, denominator)


// 	if err!=nil {
// 		fmt.Printf(err.Error())
// 	} else if remainder == 0 {
// 		fmt.Printf("The result of the integer division is %v", result)
// 	} else {
// 		fmt.Printf("The result of integer division is %v with remainder %v", result, remainder)
// 	} 

	switch {
	case err!= nil:
		fmt.Printf(err.Error())
	case remainder == 0:
		fmt.Printf("The result of the integer division %v", result)
	default:
		fmt.Printf("The result of integer division is %v with remainder %v", result, remainder)	
		
	}
}


func intDivision(numerator int, denominator int) (int,int, error) {
	var err error 
	if denominator==0 {
		err = errors.New("Cannot Divide by Zero")
		return 0, 0, err
	}
	var result int = numerator/denominator;
	var remainder int = numerator % denominator
	return result, remainder, err 
}