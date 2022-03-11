package main

import "fmt"

type subscriber struct {
	name   string
	rate   float64
	active bool
}

//print the subscriber
func main() {

	var sub1 subscriber
	sub1.name = "John Smith"
	sub1.rate = 4.99
	sub1.active = true

	var sub2 = subscriber{
		name:   "Jane Doe",
		rate:   5.99,
		active: false,
	}

	fmt.Printf("Name: %s\n", sub1.name)
	fmt.Printf("Rate: %.2f\n", sub1.rate)
	fmt.Printf("Active: %t\n", sub1.active)
	fmt.Printf("Name: %s\n", sub2.name)
	fmt.Printf("Rate: %.2f\n", sub2.rate)
	fmt.Printf("Active: %t\n", sub2.active)
}
