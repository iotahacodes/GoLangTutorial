package main

import "fmt"

func main() {
	s := "Yooo"
	fmt.Printf("%T\n", s)

	s = "changed yooo\n"
	fmt.Println(s)

	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)
}
