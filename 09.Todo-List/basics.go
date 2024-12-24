package main

import "fmt"





func basics(){
	fmt.Print("Hello Go World\n")

	
	var taskOne = "Watch Golang tut\n"
	var taskTwo = "Watch intersteller\n"
	var taskItems = []string{taskOne, taskTwo}

	printTask(taskItems)

	taskItems = addTask(taskItems, "Go for a run\n")
	taskItems = addTask(taskItems, "Go Gym")

	printTask(taskItems)
}



func printTask(taskItems []string){

	for index,task := range taskItems{
		index += 1
		fmt.Print(index,". ", task)
	}
}

func addTask(taskItems []string, newTask string) []string  {
	println()
	println("updated list")
	var updatedTaskItems =  append(taskItems, newTask)

	return updatedTaskItems
	  
}