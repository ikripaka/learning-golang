package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var counterMutex sync.Mutex

func main() {
	var counter int32
	channel := make(chan int)
	timeBefore := time.Now()
	go incrementWithMutex("<1>", channel, &counter)
	go incrementWithMutex("<2>", channel, &counter)
//	go incrementWithAtomic("<2>", channel, &counter)
//	go incrementWithAtomic("<2>", channel, &counter)

	if (<-channel + <-channel) == 2 {
		fmt.Println("Counter: ", counter, "Time:", time.Now().Sub(timeBefore))
		close(channel)
	}
}

func incrementWithMutex(processName string, channel chan int, counter *int32) {
	for i := 0; i < 1000000; i++ {
		counterMutex.Lock()
		*counter++
		counterMutex.Unlock()
		//fmt.Println(processName, "|Iteration:", i, "|Counter:", &counter)
	}
	channel <- 1
}

func incrementWithAtomic(processName string, channel chan int, counter *int32) {
	for i := 0; i < 1000000; i++ {
		atomic.AddInt32(counter, 1)
		//fmt.Println(processName, "|Iteration:", i, "|Counter:", &counter)
	}
	channel <- 1
}