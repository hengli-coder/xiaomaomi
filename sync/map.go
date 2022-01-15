package example

import (
	"fmt"
	"sync"
)

func ExampleMap() {
	var m = make(map[interface{}]interface{})
	var i interface{}
	var t, _ = (i).(int)
	fmt.Println(m[10], t)
	var sm = new(sync.Map)
	fmt.Println(sm)
	sm.Store(1, 1)
	sm.LoadOrStore(2, 2)
	sm.Load(1)
	sm.Range(nil)
	sm.LoadAndDelete(2)
	sm.Delete(1)
}
