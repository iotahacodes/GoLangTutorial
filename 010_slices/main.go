package main

import "fmt"

func main() {
	s := make([]string, 4)
	fmt.Println("s:", s)
	fmt.Println("len:", len(s))

	s[0] = "James"
	s[1] = "Lars"
	s[2] = "Kirk"
	s[3] = "Robert"

	fmt.Println("\nset:", s)

	s = append(s, "me")
	fmt.Println("\nset:", s)
	fmt.Println("len:", len(s))

	q := make([]string, len(s))
	copy(q, s)
	fmt.Println("\nCopy:", q)

	fmt.Println(s[2:])

	twoL := make([][]int, 3)
	for i := 0; i < 3; i++ {
		inner := i + 1
		twoL[i] = make([]int, inner)
		for j := 0; j < inner; j++ {
			twoL[i][j] = i + j
		}
	}
	fmt.Println("Two Layer:", twoL)
}
