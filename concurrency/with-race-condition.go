package main

import (
	"fmt"
	"sync"
)

var counterMutex sync.Mutex
var counter int

func main() {
	channel := make(chan int)
	go increment("<1>", channel)
	go increment("<2>", channel)
	if (<-channel + <-channel) == 2 {
		fmt.Println("Counter: ", counter)
		close(channel)
	}
}

func increment(processName string, channel chan int) {
	for i := 0; i < 1000000; i++ {
		counterMutex.Lock()
		counter++
		counterMutex.Unlock()
		fmt.Println(processName, "|Iteration:", i, "|Counter:", counter)
	}
	channel <- 1
}