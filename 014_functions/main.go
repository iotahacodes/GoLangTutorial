package main

import "fmt"

func mathBazi(a int, b int) (int, int) {
	return a + b, a * b
}

func main() {
	a := 85
	b := 69

	add, mult := mathBazi(a, b)
	fmt.Println(add, mult)
}
