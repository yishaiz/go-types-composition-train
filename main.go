package main

import "fmt"

// reference types (pointers, slices, maps, functions, channels)

// interface type

func main() {
	// var myMap map[string]string

	intMap := make(map[string]int)

	intMap["one"] = 1
	intMap["two"] = 2
	intMap["three"] = 3
	intMap["four"] = 4
	intMap["five"] = 5

	for key, value := range intMap {
		fmt.Println(key, value)
	}

	fmt.Println("*** delete four *** ")
	delete(intMap, "four")

	for key, value := range intMap {
		fmt.Println(key, value)
	}

	// check if exist

	el, ok := intMap["four"]
	if ok {
		fmt.Println(el, "is in map")
	} else {
		fmt.Println(el, "is NOT in map")
	}

}
