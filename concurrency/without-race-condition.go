package main

import (
	"fmt"
	//"math/rand"
	//"time"

	//"time"
)

var counter int

func main() {
	//rand.Seed(time.Now().UnixNano())
	//var waitGroup sync.WaitGroup

	//waitGroup.Add(20)
	channel := make(chan int)
	go increment("<1>", channel)
	go increment("<2>", channel)
	if (<-channel + <-channel) == 2 {
		fmt.Println("Counter: ", counter)
		close(channel)
	}
}

func increment(processName string, channel chan int ) {
	for i := 0; i < 1000000; i++ {
		counter++
		fmt.Println(processName, "|Iteration:", i, "|Counter:", counter)
	}
	fmt.Print("")
	channel <- 1
}
