package main

import "fmt"

func array() {
	var myStrings [3]string

	myStrings[0] = "cat"
	myStrings[1] = "dog"
	myStrings[2] = "fish"

	fmt.Println("First element in array is", myStrings[0])
	fmt.Println("Second element in array is", myStrings[1])
	fmt.Println("Third element in array is", myStrings[2])
}
