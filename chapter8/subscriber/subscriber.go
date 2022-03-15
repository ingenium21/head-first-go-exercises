package main

import (
	"fmt"
	"head-first-go-exercises/chapter8/magazine"
)

func printInfo(s magazine.Subscriber) {
	fmt.Printf("Name: %s\n", s.Name)
	fmt.Printf("Rate: %.2f\n", s.Rate)
	fmt.Printf("Active: %t\n", s.Active)
	fmt.Println("Home Address:", s.HomeAddress)
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
	addr1 := magazine.Address{Street: "123 Oak Street", City: "Omaha", State: "NE", PostalCode: "68111", Country: "USA"}
	sub1.HomeAddress = addr1
	printInfo(sub1)

	sub2 := defaultSubscriber("Jane Doe")
	addr2 := magazine.Address{Street: "456 Elm Street", City: "Omaha", State: "NE", PostalCode: "68111", Country: "USA"}
	sub2.HomeAddress = addr2
	printInfo(sub2)
}
