package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	foo()

	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	bar()
}

func foo() {
	fmt.Println("This is so awesome")
}

func bar() {
	fmt.Println("I am the bar function")
}
