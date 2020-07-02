package main


func Increment( counter *int32) {
	for i := 0; i < 1000000; i++ {
		*counter++
	}
}
