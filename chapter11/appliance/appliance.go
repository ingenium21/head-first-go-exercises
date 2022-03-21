package main

type Appliance interface {
	TurnOn()
	TurnOff()
}

type Fan string

func (f Fan) TurnOn() {
	println("Fan is on")
}

func (f Fan) TurnOff() {
	println("Fan is off")
}

type CoffeePot string

func (c CoffeePot) TurnOn() {
	println("Coffee pot is on")
}

func (c CoffeePot) TurnOff() {
	println("Coffee pot is off")
}

func (c CoffeePot) PourCoffee() {
	println("Pouring coffee")
}

func main() {
	//if you define a type as an interface, you can assign it to a variable of type interface
	//but you can't use functions that are not defined in the interface
	var device Appliance
	device = Fan("Windco Breeze")
	device.TurnOn()
	device.TurnOff()
	device = CoffeePot("LuxBrew")
	device.TurnOn()
	device.TurnOff()

	//if you assign it as a concrete type, you can use functions that are defined in the interface AND the concrete type
	var dev1 CoffeePot
	dev1.TurnOn()
	dev1.TurnOff()
	dev1.PourCoffee()
}
