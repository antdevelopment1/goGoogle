package main

import (
	"fmt"
	"runtime"
)

var x int
var y float64
var w int8 = 127

func main() {
	x = 42
	y = 42.73737
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(w)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)

}
