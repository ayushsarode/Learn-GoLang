package main

import (
	"fmt"

)

func main(){
	var first = "I want to learn Go\n"
	var second = "I want to attend meeting at 8\n"
	var third = "Walk at 7"

	// var taskItem = [3]string{first, second, third}

	taskItem := [3]string{first, second, third} 
	// := equivalent to var varName = ();


	// var taskItem = [2]string{first, second, third}  
	//error as arrays have fixed size, var varName = [size]variable_type


	fmt.Println("Tasks:", taskItem)
}