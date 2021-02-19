package main

import "fmt"

func main() {
	a := new(struct{})
	b := new(struct{})
	println("1: ", a, b, a == b)
	fmt.Println("------------------")
	c := new(struct{})
	d := new(struct{})
	fmt.Println("2: ", c, d, c == d)
}
