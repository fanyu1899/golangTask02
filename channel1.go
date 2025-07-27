package main

import (
	"fmt"
	"time"
)

var chan1 = make(chan int)

func put() {
	for i := range 10 {
		chan1 <- i
	}
	close(chan1)
}

func printNum() {

	for i := range chan1 {
		println(i)
	}
}

func testChannel() {
	go put()
	printNum()
}

var chan2 = make(chan int, 40)

func put2() {
	for i := range 100 {
		chan2 <- i
		fmt.Printf("put %d\n", i)
	}
	close(chan2)

}

func testChannel2() {

	go put2()

	time.Sleep(1 * time.Second)
	for i := range chan2 {
		fmt.Printf("get %d\n", i)
	}
}
