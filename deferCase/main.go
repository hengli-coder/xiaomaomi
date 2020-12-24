package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	bigSlowOperation()
}

func bigSlowOperation() {
	defer trace("big")()()
	fmt.Println("start big")
	time.Sleep(time.Second)
	fmt.Println(name{age: "10"})

	var x interface{} = 3
	var y interface{} = 4
	fmt.Println(x == y)
}

type name struct {
	age string
}

func (n *name)String() string {
	return "{  "+n.age+" }"
}

func trace(msg string) func()func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func()func() {
		t:=time.Since(start)
		log.Printf("exit %s (%s)", msg, t)
		return func() {
			log.Printf("enter func func")
		}
	}

}
