package main

import "fmt"

func main()  {
	// var intArr [3]int64
	// intArr[1] = 123
	// intArr[0] = 0124143253514364
	// intArr[2] = 144
	// fmt.Println(intArr[0])
	// fmt.Println(intArr[0:3])

	// fmt.Println(&intArr[1]) //012414325351 memory adress 
	// fmt.Println(&intArr[2])


	// intArray := [...]int32{1,2,3}
	// fmt.Println(intArray)



	// Slices:  A slice is a reference to a portion of an array. 

	// var intSlice []int32 = []int32{4,5,6}
	// fmt.Printf("The length is %v with the capacity %v", len(intSlice), cap(intSlice))
	// intSlice = append(intSlice, 7)
	// fmt.Printf("\nThe length is %v with the capacity %v\n", len(intSlice), cap(intSlice))

	// fmt.Println(intSlice)

	// //append
	// var intSlice2 []int32 = []int32{8,9}
	// intSlice = append(intSlice, intSlice2...)
	// fmt.Println((intSlice))


	//map
	var myMap = map[string]uint8{"Adam": 28, "Sarah": 45}
	fmt.Println(myMap["Adam"])
	var age, ok = myMap["Adam"]
	if !ok {
		fmt.Println("Invalid Name")
	} else{
		(fmt.Printf("The age is %v", age))
	}


	var mapArr = map[string][]string{"Name": {"Ashish", "Ajay", "Dev"}}

	fmt.Printf("Name: %v/n", mapArr["Name"][1])



	for i := 1; i < 10; i *= 2 {
		fmt.Println(i)
	}

}