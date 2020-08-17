package main

import "fmt"

var x = 10
var y int

func main() {
	a := 85
	fmt.Println(a)

	a = 69
	fmt.Println(a)

	b := 85 + 69
	fmt.Println(b)

	c := "Broooo..."
	fmt.Println(c)

	nadombe()
}

func nadombe() {
	fmt.Println(x, y)
}
