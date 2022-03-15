//Same as subscriber.go but with a pointer to the subscriber
package main

import (
	"fmt"
	"head-first-go-exercises/chapter8/magazine"
)

func printInfoPointer(s *magazine.Subscriber) {
	fmt.Printf("Name: %s\n", s.Name)
	fmt.Printf("Rate: %.2f\n", s.Rate)
	fmt.Printf("Active: %t\n", s.Active)
}

func defaultSubscriberPointer(name string) *magazine.Subscriber {
	var s magazine.Subscriber
	s.Name = name
	s.Rate = 5.99
	s.Active = true
	return &s
}

func applyDiscountPointer(s *magazine.Subscriber) {
	s.Rate = 3.99
}

func main() {
	sub1 := defaultSubscriberPointer("Aman Singh")
	//the reason why you don't need to use & is because the defaultSubscriberPointer function returns a pointer to sub1.
	applyDiscountPointer(sub1)
	printInfoPointer(sub1)

	sub2 := defaultSubscriberPointer("Beth Ryan")
	printInfoPointer(sub2)
}
