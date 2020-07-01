package main

import (
		"sync"
	"sync/atomic"
)

func IncrementWithMutex(processName string, counter *int32, waitGroup *sync.WaitGroup, counterMutex *sync.Mutex) {
	for i := 0; i < 1000000; i++ {
		counterMutex.Lock()
		*counter++
		counterMutex.Unlock()
	}
	waitGroup.Done()

}
func IncrementWithAtomic(processName string,channel chan int, counter *int32) {
	for i := 0; i < 1000000; i++ {
		atomic.AddInt32(counter, 1)
	}
	channel <- 1
}
