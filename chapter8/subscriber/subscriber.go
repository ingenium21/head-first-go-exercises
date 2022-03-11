package main

import "fmt"

type subscriber struct {
	name   string
	rate   float64
	active bool
}

func printInfo(s subscriber) {
	fmt.Printf("Name: %s\n", s.name)
	fmt.Printf("Rate: %.2f\n", s.rate)
	fmt.Printf("Active: %t\n", s.active)
}

func defaultSubscriber(name string) subscriber {
	var s subscriber
	s.name = name
	s.rate = 5.99
	s.active = true
	return s
}

//print the subscriber
func main() {

	//create two subscribers
	sub1 := defaultSubscriber("John Smith")
	sub1.rate = 4.99
	printInfo(sub1)

	sub2 := defaultSubscriber("Jane Doe")
	printInfo(sub2)
}
