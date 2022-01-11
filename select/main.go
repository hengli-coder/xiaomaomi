package main

import (
	"fmt"
	"time"
)

func main() {

	var s = make([]int, 10, 20)
	var a = s[10:14:15]
	fmt.Println(a, len(a), cap(a), len(s), cap(s))
	var c = make(chan int)
	var b = make(chan int)
	for true {
		time.Sleep(time.Second)
		select {
		case v, ok := <-c:
			fmt.Println(v, ok)
		case c <- 10:
			fmt.Println("input 10")
		case b <- 0:
			fmt.Println("input 0")
		default:
			fmt.Println("default")
		}
	}
}
