package main

import (
	"html/template"
	"log"
	"os"
)

type Invoice struct {
	Name    string
	Paid    bool
	Charges []float64
	Total   float64
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	html, err := template.ParseFiles("C:\\Users\\renat\\Documents\\GitHub\\head-first-go-exercises\\chapter16\\invoice\\bill.html")
	check(err)
	bill := Invoice{
		Name:    "Mary Gibbs",
		Paid:    true,
		Charges: []float64{23.19, 1.13, 42.79},
		Total:   67.32,
	}
	err = html.Execute(os.Stdout, bill)
	check(err)
}
