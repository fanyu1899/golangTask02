package main

import (
	"fmt"
	"time"
)

func print1() {
	for i := 0; i < 1000; i += 2 {
		fmt.Println(i)
	}
}

func print2() {
	for i := 1; i < 1000; i += 2 {
		fmt.Println(i)
	}

}
func testGoRoutine() {
	go print1()
	go print2()
	time.Sleep(1 * time.Second) // 等待 goroutine 执行

}
