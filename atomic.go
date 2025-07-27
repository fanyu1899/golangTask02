package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var (
	counter int64 = 0
	wg      sync.WaitGroup
)

func testAtomic() {
	for range 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 1000 {
				atomic.AddInt64(&counter, 1)
			}
		}()
	}

	wg.Wait()

	fmt.Println(counter)
}
