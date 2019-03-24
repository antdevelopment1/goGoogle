package main

import "fmt"

var a = "This is a string"

func main() {
	fmt.Println(a)
	var b = 32
	fmt.Printf("%T\n", b)
	// We cannot change the value type of b to a string since it proginally assined as a number
	// b = "Hello. I am tring to change the type of this value to a string even though it was a number"
	// fmt.Println(b)

}
