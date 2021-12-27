package main

import "time"

func main() {
	var c = make(chan int)
	var cb = make(chan int,3)
	go send(c)
	go receive(c)

	go send(cb)
	go receive(cb)

	time.Sleep(time.Millisecond)
}

func send(c chan int)  {
	c<-1
}

func receive(c chan int)  {
	<-c
}
