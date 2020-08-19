package main

import "fmt"

func main() {
	defer fmt.Println("Done, function returned")

	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}
}
