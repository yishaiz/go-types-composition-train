package main

import "fmt"

func array() {
	var myStrings [3]string

	myStrings[0] = "cat"
	myStrings[1] = "dog"
	myStrings[2] = "fish"

	fmt.Println("First element in array is", myStrings[0])
	fmt.Println("Second element in array is", myStrings[1])
	fmt.Println("Third element in array is", myStrings[2])
}


type Car struct {
	NumOfTires  int
	Luxury      bool
	BucketSeats bool
	Make        string
	Model       string
	Year        int
}

func main1() {
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
