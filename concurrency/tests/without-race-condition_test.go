package main

import (
	"fmt"
	"math"
	"sync"
	"testing"
	"time"
)

const MaxGoroutineNumber = 1024

func BenchmarkIncrementWithMutex(b *testing.B) {
	var counterMutex sync.Mutex
	var counter int32
	for i := 1; i <= MaxGoroutineNumber; i *= 2 {
		timeNowMinBefore := time.Now().Minute()
		timeNowSecBefore := time.Now().Second()
		b.Run(fmt.Sprint("goroutine number ", i), func(b *testing.B) {

			for j := 0; j < i; j++ {
				go IncrementWithMutex(&counter, &counterMutex)
			}
		})
		fmt.Println("(Mutex) Pass|", "goroutines number:", i, "time:", math.Abs(float64(time.Now().Minute()-timeNowMinBefore)),
			"min", math.Abs(float64(time.Now().Second()-timeNowSecBefore)), "sec")
		fmt.Println("time before:", timeNowMinBefore, "min", timeNowSecBefore, "sec", "time after:",
			time.Now().Minute(), "min", time.Now().Second(), "sec")
		fmt.Println("-- -- -- -- -- --")
	}
}

func BenchmarkIncrementWithAtomic(b *testing.B) {
	var counter int32
	channel := make(chan int)
	for i := 1; i <= MaxGoroutineNumber; i *= 2 {
		timeNowMinBefore := time.Now().Minute()
		timeNowSecBefore := time.Now().Second()
		b.Run(fmt.Sprint("goroutine number ", i), func(b *testing.B) {
			for j := 0; j < i; j++ {
				go IncrementWithAtomic(channel, &counter)
			}

		})
		fmt.Println("(Atomic) Pass|", "goroutines number:", i, "time:", math.Abs(float64(time.Now().Minute()-timeNowMinBefore)),
			"min", math.Abs(float64(time.Now().Second()-timeNowSecBefore)), "sec")
		fmt.Println("time before:", timeNowMinBefore, "min", timeNowSecBefore, "sec", "time after:",
			time.Now().Minute(), "min", time.Now().Second(), "sec")
		fmt.Println("-- -- -- -- -- --")
	}
}
