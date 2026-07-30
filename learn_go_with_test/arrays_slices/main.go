package main

// slices, in slices we dont have number limit lint in array like [5]int

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}


