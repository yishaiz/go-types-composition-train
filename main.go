package main

import "fmt"

// reference types (pointers, slices, maps, functions, channels)

// interface type

func main() {
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

// func main() {
// 	z := addTwoNumbers(2, 4)
// 	fmt.Println(z)
// }

// func addTwoNumbers(x, y int) (sum int) {
// 	sum = x + y
// 	return
// }
