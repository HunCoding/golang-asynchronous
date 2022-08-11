package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter uint64
	var wg sync.WaitGroup
	var mu sync.Mutex

	//TODO: Implement concurrency safe counter
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for c := 0; c < 1000; c++ {
				mu.Lock()
				counter++
				mu.Unlock()
			}
		}()
	}

	wg.Wait()
	fmt.Printf("counter: %v\n", counter)
}
