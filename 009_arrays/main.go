package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("a:\t", a)
	a[4] = 100
	fmt.Println("set:\t", a)
	fmt.Println("get:\t", a[4])
	fmt.Println("len:\t", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("b:\t", b)

	var twoL [2][3]int
	for i := 0; i <= 1; i++ {
		for j := 0; j <= 2; j++ {
			twoL[i][j] = i + j
		}
	}
	fmt.Println("twoL:", twoL)

}
