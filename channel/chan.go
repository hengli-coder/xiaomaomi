package main

import (
	"fmt"
	"time"
)

func main() {


	go func() {

		var nilc chan int
		nilc <- 10
	}()

	var c = make(chan int)
	var cb = make(chan int)
	//close(cb)
	go send(c)
	//go receive(c)

	//go send(cb)
	//go receive(cb)


	time.Sleep(time.Second)
	//close(c)
	close(cb)
	time.Sleep(time.Second*3)
}

func send(c chan int) {
	select {
	case c<-1:
		fmt.Println("c send data")
	//default:

	}
//	c <- 1
	fmt.Println("do not block: ",time.Now().GoString())
}

func receive(c chan int) {
	select {
	case <-c:
		fmt.Println("c is ok")
	default:
		fmt.Println("default")
	}
}
