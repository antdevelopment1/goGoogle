package main

import "fmt"

// Variables like this can be defined outside of a function
var w = 43

func numbers() {
	// Short hand variable declaration

	x := 52
	fmt.Println(x)
	// Variables assined like this have to be defined or have function blocked scope
	y := 36
	fmt.Println(y)
	z := 22
	fmt.Println(z)
	// Simple way of declaring a variable
	// Old way of declaring varbales. It if infered that it is a number Int not needed
	var h int = 8
	fmt.Println(h)

	foo()

}

func foo1() {
	fmt.Println(w)
}
