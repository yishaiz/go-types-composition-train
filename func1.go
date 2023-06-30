package main

import "fmt"

// reference types (pointers, slices, maps, functions, channels)

// interface type
  
func func1() {
	z := addTwoNumbers(2, 4)
	fmt.Println(z)
}

func addTwoNumbers(x, y int) (sum int) {
	sum = x + y
	return
}
