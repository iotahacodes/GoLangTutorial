package main

import (
	"fmt"
	"time"
)

func main() {
	today := time.Now().Weekday()
	switch today {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend")
	default:
		fmt.Println("Weekday")
	}
}
