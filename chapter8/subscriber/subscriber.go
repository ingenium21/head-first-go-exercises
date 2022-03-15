package main

import (
	"fmt"
	"head-first-go-exercises/chapter8/magazine"
)

func printInfo(s magazine.Subscriber) {
	fmt.Printf("Name: %s\n", s.Name)
	fmt.Printf("Rate: %.2f\n", s.Rate)
	fmt.Printf("Active: %t\n", s.Active)
}

func defaultSubscriber(name string) magazine.Subscriber {
	var s magazine.Subscriber
	s.Name = name
	s.Rate = 5.99
	s.Active = true
	return s
}

//change rate pointer to a new value
func applyDiscount(s *magazine.Subscriber) {
	s.Rate = 3.99
}

//print the subscriber
func main() {

	//create two subscribers
	sub1 := defaultSubscriber("John Smith")
	applyDiscount(&sub1)
	printInfo(sub1)

	sub2 := defaultSubscriber("Jane Doe")
	printInfo(sub2)
}
