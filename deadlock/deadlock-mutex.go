package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var fuel int32 = 45
var water int32 = 125

var fuelMutex sync.Mutex
var waterMutex sync.Mutex
var waitGroup sync.WaitGroup

//Situation with deadlock
func addition3(processCount int, waitGroup *sync.WaitGroup) {
	if rand.Int()%2 == 0 {
		fuelMutex.Lock()
		waterMutex.Lock()
		time.Sleep(time.Duration(time.Duration.Seconds(2)))
	} else {
		waterMutex.Lock()
		fuelMutex.Lock()
		time.Sleep(time.Duration(int64(rand.Int() * 100)))
	}

	fmt.Println("\n<", processCount, ">", "Mutex locked")
	fuel += 10
	water = fuel
	fmt.Printf("Fuel: %d, Water: %d \n", fuel, water)

	fuelMutex.Unlock()
	waterMutex.Unlock()
	fmt.Println("<", processCount, ">", "|Now unlocked")

	waitGroup.Done()
}

//Situation with 2 mutexes
func addition2(processCount int, waitGroup *sync.WaitGroup) {
	fuelMutex.Lock()
	waterMutex.Lock()

	fmt.Println("<", processCount, ">", "Mutex locked")
	fuel += 10
	water = fuel
	fmt.Printf("\n Fuel: %d, Water: %d \n", fuel, water)

	fuelMutex.Unlock()
	waterMutex.Unlock()
	fmt.Println("<", processCount, ">", "|Now unlocked")

	waitGroup.Done()
}

//Situation with double locking
func addition1(processCount int, waitGroup *sync.WaitGroup) {
	fuelMutex.Lock()
	fuelMutex.Lock()

	fmt.Println("<", processCount, ">", "Mutex locked")
	fuel += 10
	fmt.Printf("\n Array after array: %d \n", fuel)

	fuelMutex.Unlock()
	fmt.Println("<", processCount, ">", "|Now unlocked")

	waitGroup.Done()
}
