package main

import "fmt"

const a = 42
const b = 42.34556
const c = "James Bond"

// Untyped constant
// const (
// 	d = 33
// 	e = 33.333344
// 	f = "Authur Stone"
// )

// Types Constants
const (
	d int     = 66
	e float64 = 22.333
	f string  = "Hello"
)

func main() {
	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	// fmt.Printf("%T\n", a)
	// fmt.Printf("%T\n", b)
	// fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", d)
	fmt.Printf("%T\n", e)
	fmt.Printf("%T\n", f)
}
