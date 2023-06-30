package main

import "fmt"

// reference types (pointers, slices, maps, functions, channels)

// interface type
 

func func2() {
	myTotal := sumMany(2, 3, 4, 5)
	fmt.Println(myTotal)
}

 

func sumMany(nums ...int) int {
	total := 0

	for _, x := range nums {
		total += x
	}

	return total
}
 