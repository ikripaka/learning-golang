package main

import (
		"sync"
	"sync/atomic"
)

func IncrementWithMutex( counter *int32, counterMutex *sync.Mutex) {
	for i := 0; i < 1000000; i++ {
		counterMutex.Lock()
		*counter++
		counterMutex.Unlock()
	}

}
func IncrementWithAtomic(channel chan int, counter *int32) {
	for i := 0; i < 1000000; i++ {
		atomic.AddInt32(counter, 1)
	}
	channel <- 1
}
