package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
	fmt.Print("I'm Taha :) ")
	fmt.Println("I freaking love Golang!")
	folan()

	println("\nPrinting even numbers!.")
	for i := 0; i <= 20; i++ {
		if i%2 == 0 {
			fmt.Print(i)
			fmt.Print(" ")
		}
	}
	bisar()
}

func folan() {
	fmt.Println("Something from folan(). O_O")
}

func bisar() {
	fmt.Println("\nExiting the program.")
}
