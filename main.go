package main

import "fmt"

// aggregate types (array, struct)

// reference types (pointers, slices, maps, functions, channels)

// interface type

func main() {
	var myStrings [3]string

	myStrings[0] = "cat"
	myStrings[1] = "dog"
	myStrings[2] = "fish"

	fmt.Println("First element in array is", myStrings[0])
	fmt.Println("Second element in array is", myStrings[1])
	fmt.Println("Third element in array is", myStrings[2])
}
