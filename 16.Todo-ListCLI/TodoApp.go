package main

import (
	"fmt"
	"net/http"
)

var taskOne = "Watch Golang tut\n"
var taskTwo = "Watch intersteller\n"
var taskThree = "Read a book for 15 min\n"
var taskItems = []string{taskOne, taskTwo, taskThree}

func main()  {
	
	fmt.Print("Sever is running at 8000")

	http.HandleFunc("/", helloUser)

	http.HandleFunc("/show-tasks", showTasks)


	http.ListenAndServe(":8000", nil)

}

func showTasks(writer http.ResponseWriter, request *http.Request) {
	for _, task := range taskItems {
		fmt.Fprintln(writer, task)
	}
	
}

func helloUser(writer http.ResponseWriter,request *http.Request)  {

	Check := "Hello, is Im workin"
	fmt.Fprintln(writer, Check)
}