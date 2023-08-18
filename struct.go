package main

import "fmt"


type Car1 struct {
	NumOfTires  int
	Luxury      bool
	BucketSeats bool
	Make        string
	Model       string
	Year        int
}

func struct1() {
	myCar := Car1{
		NumOfTires:  4,
		Luxury:      true,
		BucketSeats: true,
		Make:        "Volco",
		Model:       "XC90",
		Year:        2019,
	}

	fmt.Printf("My car is a %d %s %s", myCar.Year, myCar.Make, myCar.Model)

}


 