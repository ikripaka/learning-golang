package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAddition3(test *testing.T) {
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
}

func TestAddition1(test *testing.T) {
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
}

func TestAddition2(test *testing.T) {
	fmt.Printf("Arrays before fuel: %d, water: %d\n", fuel, water)

	waitGroup.Add(10)
	for i := 0; i < 10; i++ {
		go addition2(i, &waitGroup)
	}

	waitGroup.Wait()
	fmt.Printf("All is done! Water: %d, Fuel: %d", water, fuel)
}
