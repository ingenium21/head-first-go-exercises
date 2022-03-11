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

func showInfo(p part) {
	fmt.Printf("Description: %s\n", p.description)
	fmt.Printf("Count: %d\n", p.count)
	fmt.Printf("Price: %.2f\n", p.price)
}

func minimimumOrder(description string, price float64) part {
	var p part
	p.description = description
	p.count = 100
	p.price = price
	return p
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

	bolts1 := minimimumOrder(bolts.description, bolts.price)

	showInfo(bolts)
	showInfo(bolts1)
}
