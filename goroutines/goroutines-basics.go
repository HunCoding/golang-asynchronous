package main

import (
	"fmt"
	"time"
)

func fun(value string) {
	for i := 0; i < 3; i++ {
		fmt.Println(value)
		time.Sleep(1 * time.Millisecond)
	}
}

func test() {
	// Direct call
	fun("Direct call")

	//write goroutines with differents variantes for function call
	fgx := fun
	go fgx("goroutine - 2")

	//goroutine function call
	go fun("goroutine - 1")

	// goroutine with anonymous value call
	go func() {
		fun("goroutine - 3")
	}()

	// wait for goroutine to end
	time.Sleep(5 * time.Millisecond)

	fmt.Println("Done...")
}
