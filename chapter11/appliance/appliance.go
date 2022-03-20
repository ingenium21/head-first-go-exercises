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

func main() {
	var device Appliance
	device = Fan("Windco Breeze")
	device.TurnOn()
	device.TurnOff()
	device = CoffeePot("LuxBrew")
	device.TurnOn()
	device.TurnOff()

}
