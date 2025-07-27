package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	mu    sync.Mutex
	count int
)

func increment() {
	for range 1000 {
		mu.Lock()
		count++
		mu.Unlock()
	}
}

func testMutex() {
	for range 10 {
		go increment()
	}

	time.Sleep(4 * time.Second)

	fmt.Println(count)
}
