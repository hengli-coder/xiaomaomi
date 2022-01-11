package main

import (
	"fmt"
	"time"
)

func main() {
	time.NewTimer(0)
	for true {
		fmt.Println("do again")
		time.Sleep(time.Second)
	}
}
