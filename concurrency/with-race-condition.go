package main

import (
	"fmt"
	//"math/rand"
	//"time"

	//"time"
)


func main() {
	var counter int
	channel := make(chan int)
	go increment("<1>", channel, &counter)
	go increment("<2>", channel, &counter)
	if (<-channel + <-channel) == 2 {
		fmt.Println("Counter: ", counter)
		close(channel)
	}
}

func increment(processName string, channel chan int, counter *int) {
	for i := 0; i < 1000000; i++ {
		*counter++
		fmt.Println(processName, "|Iteration:", i, "|Counter:", &counter)
	}
	fmt.Print("")
	channel <- 1
}
