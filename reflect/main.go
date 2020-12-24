package main

import (
	"fmt"
	"reflect"
)

func main()  {
	var i interface{}
	i = 3
	t:=reflect.TypeOf(i)
	fmt.Println(t)
	val:=reflect.ValueOf(i)
	fmt.Println(val)
	fmt.Printf("%v\n",val)
	fmt.Println(val.String())
}
