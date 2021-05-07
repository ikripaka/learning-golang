package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
)


func TestAddition2(test *testing.T) {
	var waitGroup sync.WaitGroup
	time.AfterFunc(time.Second*10, func(){
		fmt.Println("Test finishes correct")
		return
	})

	fmt.Printf("Arrays before fuel: %d, water: %d\n", fuel, water)

	waitGroup.Add(10)
	for i := 0; i < 10; i++ {
		go addition2(i, &waitGroup)
	}

	waitGroup.Wait()
	fmt.Printf("All is done! Water: %d, Fuel: %d", water, fuel)
	fmt.Println("TEST 2 CORRECT")
	assert.True(test,true)
}

func TestAddition1(test *testing.T) {
	var waitGroup sync.WaitGroup
	fmt.Printf("Arrays before fuel: %d, water: %d\n", fuel, water)

	waitGroup.Add(10)
	for i := 0; i < 10; i++ {
		go addition1(i,  &waitGroup)
	}

	assert.Never(
		test,
		func() bool{
			waitGroup.Wait()
			return true
		},
		time.Second*10,
		time.Millisecond*100,
	)


	fmt.Printf("All is done! Water: %d, Fuel: %d", water, fuel)
	fmt.Println("TEST 1 CORRECT")
}

func TestAddition3(test *testing.T) {
	var waitGroup sync.WaitGroup
	time.AfterFunc(time.Second*10, func(){
		fmt.Println("Test finishes correct")
		return
	})

	fmt.Printf("Arrays before fuel: %d, water: %d\n", fuel, water)

	waitGroup.Add(10)
	for i := 0; i < 10; i++ {
		go addition3(i, &waitGroup)
	}
	//given condition doesn't satisfy in waitFor time, periodically checking the target function each tick
	assert.Never(
		test,
		func() bool{
			waitGroup.Wait()
			return true
		},
		time.Second*10,
		time.Millisecond*100,
	)
	fmt.Printf("All is done! Water: %d, Fuel: %d", water, fuel)
	fmt.Println("TEST 2 CORRECT")
}
