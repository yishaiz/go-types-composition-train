package main

import (
	"fmt"
	"sort"
)

// reference types (pointers, slices, maps, functions, channels)

// interface type

func slices() {
	// var animals []string

	animals := []string{}
	animals = append(animals, "dog")
	animals = append(animals, "fish")
	animals = append(animals, "cat")
	animals = append(animals, "horse")

	fmt.Println(animals)

	for i, x := range animals {
		fmt.Println(i, x)
	}

	fmt.Println("Element 0 is", animals[0])

	fmt.Println("First 2 elements are", animals[0:2])

	fmt.Println("The slice is", len(animals), "elements long")

	// sort
	fmt.Println("Is is sorted?", sort.StringsAreSorted(animals))

	sort.Strings(animals)

	fmt.Println("Is is sorted?", sort.StringsAreSorted(animals))

	fmt.Println(animals)

	animals = DeleteFromSlice(animals, 1)
}

func DeleteFromSlice(a []string, i int) []string {
	fmt.Println("debug")
	a[i] = a[len(a)-1]
	fmt.Println(a)
	a[len(a)-1] = ""
	fmt.Println(a)
	a = a[:len(a)-1]
	fmt.Println(a)

	return a
}
