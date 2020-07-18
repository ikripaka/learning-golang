package main

import (
	"fmt"
	"time"
	//"math/rand"
	//"time"
	//"time"
)

func main() {
	var counter int
	channel := make(chan int)
	timeBefore := time.Now()
	counter++
	go increment("<1>", channel, &counter)
	go increment("<2>", channel, &counter)
	if (<-channel + <-channel) == 2 {
		fmt.Println("Counter: ", counter, "Time:", time.Now().Sub(timeBefore))
		close(channel)
	}
}

func increment(processName string, channel chan int, counter *int) {
	for i := 0; i < 250000000; i++ {
		*counter++
		//fmt.Println(processName, "|Iteration:", i, "|Counter:", *counter)
	}
	channel <- 1
}
