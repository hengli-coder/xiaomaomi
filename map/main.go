package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

const (
	a = iota
	b
	c
)

func main()  {

	fmt.Println(a,b,c)

	var b []int
	b = append(b, 1,2,3)

	var m = make(map[int]int)
	m[0]=1
	fmt.Println(m[0])

	var s1 string
	var s2 string

	var s3 = s1+s2

	fmt.Printf("%p\n",&s1)
	fmt.Printf("%p\n",&s3)

	fmt.Println(&s3==&s2)

	fmt.Println((*reflect.StringHeader)(unsafe.Pointer(&s3)).Data,(*reflect.StringHeader)(unsafe.Pointer(&s1)).Data)
}
