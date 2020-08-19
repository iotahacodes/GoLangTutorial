package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func main() {
	a := 85
	b := 69

	fmt.Println(plus(a, b))
}
