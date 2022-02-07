package main

import (
	"fmt"
	"reflect"
)

func main() {
	//declaring variable that holds a pointer to an int
	var myInt int = 5
	var myIntPointer *int
	myIntPointer = &myInt
	fmt.Println("myIntPointer:", myIntPointer)
	fmt.Println("myIntPointer type:", reflect.TypeOf(myIntPointer))
	fmt.Println("myIntPointer value:", *myIntPointer)

	var greeting string = "Hello, Go!"
	fmt.Println(greeting)
	fmt.Println(&greeting)
	fmt.Println(reflect.TypeOf(&greeting))
}
