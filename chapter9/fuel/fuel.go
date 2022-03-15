package main

import "fmt"

type gallons float64
type liters float64
type mililiters float64

func (l liters) toGallons() gallons {
	return gallons(l * 0.264)
}

func (m mililiters) toGallons() gallons {
	return gallons(m * 0.000264)
}

func (g gallons) toLiters() liters {
	return liters(g * 3.785)
}

func (g gallons) toMililiters() mililiters {
	return mililiters(g * 3785.41)
}

func main() {
	soda := liters(2)
	fmt.Printf("%0.3f liters is %0.3f gallons\n", soda, soda.toGallons())
	water := mililiters(500)
	fmt.Printf("%0.3f mililiters is %0.3f gallons\n", water, water.toGallons())
	milk := gallons(2)
	fmt.Printf("%0.3f gallons is %0.3f liters\n", milk, milk.toLiters())
	fmt.Printf("%0.3f gallons is %0.3f mililiters\n", milk, milk.toMililiters())
}
