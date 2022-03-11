package main

import "fmt"

type part struct {
	description string
	count       int
	price       float64
}

type car struct {
	name     string
	topSpeed int
}

func main() {
	//two ways to declare a type struct
	var porsche = car{
		name:     "Porsche 911 R",
		topSpeed: 323,
	}
	fmt.Printf("Name: %s\n", porsche.name)
	fmt.Printf("Top Speed: %d\n", porsche.topSpeed)

	var bolts part
	bolts.description = "Hex Bolts"
	bolts.count = 24
	bolts.price = 3.99
	fmt.Printf("Description: %s\n", bolts.description)
	fmt.Printf("Count: %d\n", bolts.count)
	fmt.Printf("Price: %.2f\n", bolts.price)

}
