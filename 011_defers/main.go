package main

import "fmt"

func main() {
	defer fmt.Println("\nDone, function returned")

	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	fmt.Println()
	lastInFirstOut()
}

func lastInFirstOut() {
	fmt.Println("Counting")

	// Deferred function calls are pushed onto a stack. When a function returns
	// its deferred calls are executed in last-in-first-out order.
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("lastInFirstOut() returned")
}
