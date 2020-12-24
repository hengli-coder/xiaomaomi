package main

import (
	"fmt"
	"time"
)

/*
1、编译器可能会认为打印和赋值是不相关的操作，内部优化代码，交换执行顺序
2、不同的CPU之间有属于自己得临时缓存，在并发执行的时候，因为没有使用通道或者互斥锁进行内存同步，不同CPU之间的临时缓存不可见，会出现语句不一致的问题
读写锁的问题同样如此，在读取数据的时候加锁，是为了同步CPU的缓存，使得之前的修改，能同时被所有CPU看到，同时，为了保证读取数据的院子操作不会被干扰和插入，也因此，读锁会隔离写锁的请求。
*/
var x,y int

func main()  {
	go func() {
		x = 1
		fmt.Println("y : ",y)
	}()

	go func() {
		y = 1
		fmt.Println("x : ",x)
	}()

	time.Sleep(time.Second)
}
