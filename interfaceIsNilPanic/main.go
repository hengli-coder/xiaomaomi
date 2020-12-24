package main

import (
	"fmt"
	"io"
)

func main() {
	var b name
	//b = new(bytes.Buffer)
	f(b)
	var ch chan name
	var ch2 chan name
	fmt.Println(ch==ch2)
}

func f(writer io.Writer)  {
	if writer != nil{
		writer.Write([]byte("done\n"))
	}
}

type name struct {
	age int
	s []int
}

func (n2 name) Write(p []byte) (n int, err error) {
	fmt.Println("zhe shi ce shi")
	return
}
