package main

import (
	"fmt"
	"math"
	"time"
)

func main() {

	//var a = math.MaxInt32
	var b = math.MaxInt
	fmt.Printf("%0b\n",b)
	fmt.Println(len("111111111111111111111111111111111111111111111111111111111111111"))
	time.NewTimer(0)
	for true {
		fmt.Println("do again")
		time.Sleep(time.Second)
	}
}
