package main

import "fmt"

//编译器优化问题

func main() {
	var a = struct{}{}
	var b = struct{}{}
	fmt.Printf("%p,%p,%t\n", &a, &b, &a == &b)

	var m = new(struct{})
	var n = new(struct{})
	fmt.Println(&m, &n, &m == &n)
}
