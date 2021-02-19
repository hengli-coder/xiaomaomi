package main

import "fmt"

var s = new(struct{})

func main() {
	println("3: ", s)
	a := new(struct{})
	b := new(struct{})
	println("1: ", a, b, a == b)
	fmt.Println("------------------")
	c := new(struct{})
	d := new(struct{})
	fmt.Println("2: ", c, d, c == d)
	println("5: ", s, c, d, s == c)
	print()
}

//go:noinline
func print() {
	a := new(struct{})
	b := new(struct{})
	println("4: ", a, b, a == b)
}
