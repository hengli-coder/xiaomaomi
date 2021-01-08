package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

type Memo struct {
	f     Func
	mu    sync.Mutex
	cache map[string]*entry
}

func NewMemo(f Func) *Memo {
	return &Memo{f: f, cache: map[string]*entry{}}
}

type Func func(key string) (interface{}, error)

type entry struct {
	count int
	r     result
	ready chan struct{}
}

func newEntry() *entry {
	return &entry{ready: make(chan struct{})}
}

type result struct {
	value interface{}
	err   error
}

func (m *Memo) Get(key string) (interface{}, error) {

	m.mu.Lock()
	res, ok := m.cache[key]
	if !ok {
		res = newEntry()
		m.cache[key] = res
		res.count++
		m.mu.Unlock()
		//可以在这里添加重试
		res.r.value, res.r.err = m.f(key)
		close(res.ready)
	} else {
		res.count++
		m.mu.Unlock()
		<-res.ready
	}

	//如果count取值为0则需要删除相关索引
	m.mu.Lock()
	if res.count == 0 {
		delete(m.cache, key)
	}
	m.mu.Unlock()
	return res.r.value, res.r.err
}

func httpGetNBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func main() {
	m := NewMemo(httpGetNBody)
	var urls = []string{"www.baidu.com", "www.jd.com"}
	for _, url := range urls {
		start := time.Now()
		val, err := m.Get(url)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("%s %s %d bytes\n", url, time.Since(start), len(val.([]byte)))
	}
}

type request struct {
	key      string
	response chan result
}

type MemoChan struct {
	r chan request
}

func NewMemoChan(f Func) *MemoChan {
	m := &MemoChan{r: make(chan request)}
	go m.serve(f)
	return m
}

func (m *MemoChan) GetFromChan(key string) (interface{}, error) {
	r := request{response: make(chan result), key: key}
	m.r <- r
	result := <-r.response
	return result.value, result.err
}

//没有删除对应的索引，没有赋值，会导致崩溃，如何删除
func (m *MemoChan) serve(f Func) {
	cache := make(map[string]*entry)
	//需要在入队列之前就做判断
	for r := range m.r {
		e := cache[r.key]
		if e == nil {
			e = newEntry()
			cache[r.key] = e
			go e.call(f, r.key)
		}
		//计数做统计，便于后面删除索引
		e.count++
		go e.delivery(r.response)

	}
}

//可以在这里做重试
func (e *entry) call(f Func, url string) {
	e.r.value, e.r.err = f(url)
	close(e.ready)
}

func (e *entry) delivery(response chan<- result) {
	<-e.ready
	response <- e.r
	close(response)
	//并发减，需要加锁，当数量减为一的时候需要，删除对应的索引
	e.count--
	if e.count == 0 {

	}
}
