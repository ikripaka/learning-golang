package main

import (
	"fmt"
	"sync"
	"testing"
)

func benchmarkIncrementWithMutex(b *testing.B, counterMutex *sync.Mutex, counter *int32, waitGroup *sync.WaitGroup) {

	for i := 0; i < b.N; i++ {
		go IncrementWithMutex(fmt.Sprint("<", i, ">"), counter, waitGroup, counterMutex)
	}
	waitGroup.Wait()

	fmt.Println("(Mutex) Pass")

}
func BenchmarkIncrementWithMutex(b *testing.B) {
	var counterMutex sync.Mutex
	var counter int32
	var waitGroup sync.WaitGroup
	for i := 1; i < b.N; i *= 2 {
		waitGroup.Add(i)
		benchmarkIncrementWithMutex(b, &counterMutex, &counter, &waitGroup)
	}
}

func BenchmarkIncrementWithAtomic(b *testing.B) {
	var counter int32
	channel := make(chan int)
	for i := 0; i < b.N; i++ {
		go IncrementWithAtomic(fmt.Sprint("<", i, ">"), channel, &counter)
	}
	if (<-channel + <-channel) == 2 {
		fmt.Println("(Atomic) Pass")
	}
}
