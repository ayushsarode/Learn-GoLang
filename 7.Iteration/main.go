package main

import "fmt"

func main()  {
	var first = "I want to learn Go\n"
	var second = "I want to attend meeting at 8\n"
	var third = "Walk at 7"

	taskItems := [3]string{first, second, third} 

	// for _,task := range taskItems {
	// 	fmt.Print(task)
	// }

	//the _ is referred to as the blank identifier. It is used to ignore a value returned by a function or produced by an iterator when it’s not needed.


	for index,task := range taskItems {
		// fmt.Print(index + 1," ",task)
		fmt.Printf("%d. %s\n", index+1, task)
	}
}