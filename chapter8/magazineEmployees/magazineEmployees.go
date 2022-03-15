package main

import (
	"fmt"
	"head-first-go-exercises/chapter8/magazine"
)

func main() {
	var employee magazine.Employee
	employee.Name = "Joy Carr"
	employee.Salary = 60000
	employee.Street = "Avenida Carlos Salazar 127"
	employee.City = "H. Matamoros"
	employee.State = "Tamaulipas"
	employee.PostalCode = "87390"
	fmt.Printf("Name: %s\n", employee.Name)
	fmt.Printf("Salary: %.2f\n", employee.Salary)
	fmt.Printf("Home Address: %s\n", employee.Address)
}
