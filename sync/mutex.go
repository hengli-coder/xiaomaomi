package example

import (
	"fmt"
	"sync"
)

func or()  {
	fmt.Println(1&(1|4))
}

/*
golang 中的锁是不可重入的
*/

func ExampleMutex() {
	retry(sync.Mutex{}, 10)
}

func retry(mutex sync.Mutex, count int) {
	mutex.Lock()
	defer mutex.Unlock()
	fmt.Println(count)
	count--
	if count == 0 {
		return
	}

	retry(mutex, count)
}
