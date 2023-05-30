package main

import (
	"fmt"
	"sync"
)

func main() {
	count := 20
	fmt.Printf("Going to create %d goroutines.\n", count)

	//flag := true  //DeadLock
	flag := false

	var waitGroup sync.WaitGroup

	fmt.Printf("%#v\n", waitGroup)
	for i := 0; i < count; i++ {
		waitGroup.Add(1)
		go func(x int) {
			defer waitGroup.Done()
			fmt.Printf("%d ", x)
		}(i)
	}

	// More Add() calls
	if flag {
		waitGroup.Add(1)
	} else {
		// More Done() calls
		waitGroup.Done()
	}

	fmt.Printf("%#v\n", waitGroup)
	waitGroup.Wait()
	fmt.Println("\nExiting...")
}
