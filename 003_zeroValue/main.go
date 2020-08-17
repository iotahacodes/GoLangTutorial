package main

import "fmt"

var s string
var n int

func main() {
	fmt.Println("First", s, "Third") // you can see that there's an empty space in the middle.
	fmt.Printf("%T\n", s)

	s = "Second"
	fmt.Println("First", s, "Third")

	fmt.Println()
	fmt.Println(n) // now you can see that the value of n is equals to 0
	fmt.Printf("%T\n", n)
}
