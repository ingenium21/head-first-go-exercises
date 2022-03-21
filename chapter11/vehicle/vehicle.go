package main

import "fmt"

type Car string

func (c Car) Accelerate() {
	fmt.Println("Accelerating")
}
func (c Car) Brake() {
	fmt.Println("Stopping")
}

func (c Car) steer(direction string) {
	fmt.Println("Turning", direction)
}

type Truck string

func (t Truck) Accelerate() {
	fmt.Println("Speeding up")
}
func (t Truck) Brake() {
	fmt.Println("Stopping")
}
func (t Truck) Steer(direction string) {
	fmt.Println("Turning", direction)
}
func (t Truck) loadCargo(cargo string) {
	fmt.Println("Loading", cargo)
}

type Vehicle interface {
	Accelerate()
	Brake()
	Steer(string)
}

func TryVehicle(vehicle Vehicle) {
	vehicle.Accelerate()
	vehicle.Steer("left")
	vehicle.Steer("right")
	vehicle.Brake()
	truck, ok := vehicle.(Truck)
	if ok {
		truck.loadCargo("test cargo")
	} else {
		fmt.Println("Not a truck")
	}

}

func main() {
	TryVehicle(Truck("Ford F150"))

}
