package main

import "fmt"

// reference types (pointers, slices, maps, functions, channels)

// interface type

// type Animal struct {

type Animal interface {
	Says() string
	HowManyLegs() int
}

type Dog struct {
	Name         string
	Sound        string
	NumberOfLegs int
}

func (d *Dog) Says() string {
	return d.Sound
}

func (d *Dog) HowManyLegs() int {
	return d.NumberOfLegs
}

type Cat struct {
	Name         string
	Sound        string
	NumberOfLegs int
	HasTail      bool
}

func (c *Cat) Says() string {
	return c.Sound
}

func (c *Cat) HowManyLegs() int {
	return c.NumberOfLegs
}

func main() {
	dog := Dog{
		Name:         "dog",
		Sound:        "woof",
		NumberOfLegs: 4,
	}

	Riddle(&dog)

	fmt.Println()

	var cat Cat
	cat.Name = "cat"
	cat.NumberOfLegs = 4
	cat.Sound = "meow"
	cat.HasTail = true

	Riddle(&cat)

}

func Riddle(a Animal) {
	riddle := fmt.Sprintf(`This animal says "%s" and has %d legs. What animal is it ?`, a.Says(), a.HowManyLegs())

	fmt.Print(riddle)
}
