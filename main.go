package main

import "fmt"

// aggregate types (array, struct)

// reference types (pointers, slices, maps, functions, channels)

// interface type

type Car struct {
	NumOfTires  int
	Luxury      bool
	BucketSeats bool
	Make        string
	Model       string
	Year        int
}

func main() {
	myCar := Car{
		NumOfTires:  4,
		Luxury:      true,
		BucketSeats: true,
		Make:        "Volco",
		Model:       "XC90",
		Year:        2019,
	}

	fmt.Printf("My car is a %d %s %s", myCar.Year, myCar.Make, myCar.Model)

}
