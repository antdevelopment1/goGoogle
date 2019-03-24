package main

import "fmt"

type dog string
type cat string
type number int
type limit int

var poodle dog
var kitten cat = "kitten"
var integer number
var max limit

func main() {
	poodle = "Poodle"
	fmt.Println(poodle)
	fmt.Println(kitten)
	integer = 100
	fmt.Println(integer)
	max = 5000
	fmt.Println(max)
}
