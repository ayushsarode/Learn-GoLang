package main

import ("fmt")

//declaring const 
	const Value string = "efn3qogeovnv" //public




func main() {
	fmt.Println("Variables")
	var username string = "ayush"
	fmt.Println(username)
	fmt.Printf("Variable is the type of: %T", username)


	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Type of the Variable: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)



	//float32 gives 5 digits after point
	// var smallFloat float32 = 255.236374735737
	var smallFloat float64 = 255.266264634646
	fmt.Println(smallFloat)
	fmt.Printf("Type of variable: %T \n", smallFloat)



	//default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Type of variable: %T ", anotherVariable)


	//in this its nothin blank 
	var defaultStringVariable string
	fmt.Println(defaultStringVariable)
	fmt.Printf("Type %T \n", defaultStringVariable)


	// implicit type 
	var website = "www.google.com"
	fmt.Println(website)

	// no var style 
	
	numberOfUser:= 30000.0
	fmt.Println(numberOfUser)

	// there is a way to opt-out of declaring the type of a variable by using the walrus operator
	// :=

	// requirement: must be used in a function (main counts)
	// must assign a value right away, as the type is implied by the assigned value

	fmt.Println(Value)




	maths()
}

