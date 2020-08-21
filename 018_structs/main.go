package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)
	fmt.Println(s.age)
}
