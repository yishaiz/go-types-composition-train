package main

import "fmt"


func main2() {
	x := 10
	myFirstPointer := &x

	fmt.Println("x is", x)
	fmt.Println("myFirstPointer is", myFirstPointer)
	fmt.Println("myFirstPointer is", *myFirstPointer)

	*myFirstPointer = 15

	fmt.Println("x is now", x)
	fmt.Println("myFirstPointer is now", *myFirstPointer)

	changeValueOfPointer(&x)
	fmt.Println("After function call, x is now", x)
}

func changeValueOfPointer(num *int) {
	*num = 25
}