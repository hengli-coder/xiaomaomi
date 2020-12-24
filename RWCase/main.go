package main

import (
	"fmt"
	"sync"
	"time"
)

var i = 0
var rw sync.RWMutex
var mu sync.Mutex

func main()  {
	MuCase()
	start:=time.Now()
	 go Get()
	 fmt.Println("------------------------")
	time.Sleep(time.Second)
	 fmt.Println(time.Since(start))
	 go Set(8)
	 time.Sleep(time.Second*10)
}

func Get() int {
	rw.RLock()
	fmt.Println("拿到读锁")
	defer rw.RUnlock()
	time.Sleep(time.Second*8)
	return i
}

func Set(n int) int {
	fmt.Println("开始设置值")
	start:=time.Now()
	rw.Lock()
	fmt.Println("拿到写锁")
	defer rw.Unlock()
	defer fmt.Println(time.Since(start))
	i = n
	return i
}

func MuCase()  {
	var a int
	rw.RLock()
	a = i
	mu.Lock()
	a = 3
	mu.Unlock()
	rw.Unlock()
	fmt.Println(a)
}

