package main

import (
	"fmt"
	"sync"
	"testing"
)

func BenchmarkIncrementWithMutex(b *testing.B) {
	var counterMutex sync.Mutex
	var counter int32
	channel := make(chan int)
	go IncrementWithMutex("<1>", &counter, channel, &counterMutex)
	go IncrementWithMutex("<2>", &counter,channel, &counterMutex)
	if(<-channel + <-channel) == 2{
		fmt.Println("(Mutex) Pass")
	}
}
func BenchmarkIncrementWithAtomic(b *testing.B) {
	var counter int32
	channel := make(chan int)
	go IncrementWithAtomic("<1>", channel, &counter)
	go IncrementWithAtomic("<2>", channel,&counter)
	if(<-channel + <-channel) == 2{
		fmt.Println("(Atomic) Pass")
	}
}
