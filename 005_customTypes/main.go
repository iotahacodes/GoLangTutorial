package main

import "fmt"

var x string

type jafar string

var y jafar

func main() {
	y = "I'm literally in a jafar ://"
	x = string(y) + "\nnow i (x) am in jafar too :)"
	fmt.Println(x)
}
