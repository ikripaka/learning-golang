package main

import (
	"fmt"
	"sync"
)

var counterMutex sync.Mutex

func main() {
	var counter int
	channel := make(chan int)
	go incrementWithMutex("<1>", channel, &counter)
	go incrementWithMutex("<2>", channel, &counter)
	if (<-channel + <-channel) == 2 {
		fmt.Println("Counter: ", counter)
		close(channel)
	}
}

func incrementWithMutex(processName string, channel chan int, counter *int) {
	for i := 0; i < 1000000; i++ {
		counterMutex.Lock()
		*counter++
		counterMutex.Unlock()
		fmt.Println(processName, "|Iteration:", i, "|Counter:", &counter)
	}
	channel <- 1
}