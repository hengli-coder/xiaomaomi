package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
为了能保证操作和返回结果的一一对应，需要将返回结果的channel和请求的操作一一绑定，且不会因为传输复用channel，导致阻塞
*/

var req chan request

type request struct {
	opt  string
	resp chan interface{}
}

func main() {
	r := request{opt: "Get", resp: make(chan interface{})}
	Send(r)
	Obtain(r)
}

func init() {
	go OptMap()
}

type cacheVal struct {
	b        []byte
	describe string
	ready    chan struct{}
}

func OptMap() {
	cache := make(map[string]*cacheVal)
	for s := range req {
		c, ok := cache[s.opt]
		if !ok {
			c = &cacheVal{b: nil, describe: "", ready: make(chan struct{})}
			//do something to set value.使用串行操作太阻塞效率了，大家都在等所以改动过其中一部分进行并行操作
			go func() {
				b, err := HttpGet(s.opt)
				if err != nil {
					fmt.Println(err)
				}

				c.b = b
				c.describe = "this is cache test"
				close(c.ready)
			}()
		} else {
			//ready channel实现了串行化操作，将同样的请求串行化为一个请求多次重试。减少无效请求的数量
			//同样会阻塞后续的请求，因此最好放到一个goroutines中，因此正好放到deliver中。
			//<-c.ready
		}

		go Deliver(c, s.resp)
	}

}

func Send(reqObj request) {
	req <- reqObj
}

func Deliver(result *cacheVal, resp chan<- interface{}) {

	//从一个关闭的通道中获取数据，即是获取一个零值，且不会阻塞。但是每次都会有一个冗余的结构体零值出现。如何避免呢
	<-result.ready
	resp <- result
	close(resp)
}

func Obtain(req request) (interface{}, error) {
	result := <-req.resp
	if result == nil {
		return nil, nil
	}

	return result, nil
}

func HttpGet(opt string) ([]byte, error) {
	resp, err := http.Get(opt)
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	_ = resp.Body.Close()
	return b, nil
}
