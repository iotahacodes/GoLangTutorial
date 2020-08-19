package main

import "fmt"

func mathBazi(nums ...int) (int, int) {
	add := 0
	for _, num := range nums {
		add += num
	}

	mul := 1
	for _, num := range nums {
		mul *= num
	}

	return add, mul
}

func main() {
	add, mul := mathBazi(1, 2, 3, 7)
	fmt.Println(add, mul)
}
