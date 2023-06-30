package main

import "fmt"

// reference types (pointers, slices, maps, functions, channels)

// interface type

type Animal1 struct {
	Name         string
	Sound        string
	NumberOfLegs int
}

func (a *Animal1) Says() {
	fmt.Printf("A %s says %s", a.Name, a.Sound)
	fmt.Println()
}

func (a *Animal1) HowManyLegs() {
	fmt.Printf("A %s has %d legs", a.Name, a.NumberOfLegs)
	fmt.Println()
}

func main1() {
	var dog Animal1
	dog.Name = "dog"
	dog.Sound = "woof"
	dog.NumberOfLegs = 4

	dog.Says()
	dog.HowManyLegs() 

	cat := Animal1{
		Name:         "cat",
		Sound:        "meow",
		NumberOfLegs: 4,
	}

	cat.Says()
	cat.HowManyLegs() 
}
