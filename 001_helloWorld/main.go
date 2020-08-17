package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
	fmt.Print("I'm Taha :) ")
	fmt.Println("I freaking love Golang!")
	folan()

	for i := 0; i <= 10; i++ {
		fmt.Print(i)
		fmt.Print(" ")
	}
	println("Done Counting.")
}

func folan() {
	fmt.Println("Something from folan(). O_O")
}
