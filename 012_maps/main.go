package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["k1"] = 85
	m["k2"] = 69

	fmt.Println("map:\t", m)
	fmt.Println("len:\t", len(m))

	delete(m, "k2")
	fmt.Println("map:\t", m)

	tmp := map[string]int{"jafar": 01, "ghasem": 02}
	fmt.Println(tmp)
}
