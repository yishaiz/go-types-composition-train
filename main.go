package main

import "fmt"

// reference types (pointers, slices, maps, functions, channels)

// interface type

type Dog struct {
	Name         string
	Sound        string
	NumberOfLegs int
}

type Cat struct {
	Name         string
	Sound        string
	NumberOfLegs int
	HasTail      bool
}

func main() {
	dog := Dog{
		Name:         "dog",
		Sound:        "woof",
		NumberOfLegs: 4,
	}

	Riddle(dog)

	var cat Cat
	cat.Name = "cat"
	cat.NumberOfLegs = 4
	cat.Sound = "meow"
	cat.HasTail = true

	Riddle(cat)

}

func Riddle(d Dog) {
	riddle := fmt.Sprintf(`This animal says "%s" and has %d legs. What animal is it ?`, d.Sound, d.NumberOfLegs)

	fmt.Print(riddle)
}
