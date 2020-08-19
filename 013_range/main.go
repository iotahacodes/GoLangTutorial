package main

import "fmt"

func main() {
	num := []int{1, 2, 3}
	sum := 0

	for _, num := range num {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range num {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kv := map[string]string{"j": "jafar", "n": "naghi", "s": "soosan"}
	for k, v := range kv {
		fmt.Printf("%s --> %s\n", k, v)
	}

	for k := range kv {
		fmt.Println("Keys:", k)
	}

	// range on strings iterates over Unicode code points.
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
