package main

import "sync"

type Element struct{
	mutex *sync.Mutex
	data *string
}