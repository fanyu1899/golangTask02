package main

import "fmt"

func add(p *int) {
	*p = *p + 10
}

func updateSlice(slice []int) {
	for i := 0; i < len(slice); i++ {
		slice[i] = slice[i] * 2
	}
}

func testArray(arr [10]int) {
	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] * 2
	}
}

func testArrayPoint(arr *[10]int) {
	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] * 2
	}
}

func testPoint() {
	a := 1
	add(&a)
	fmt.Println(a)

	slice := make([]int, 10)
	var arr [10]int
	for i := 0; i < 10; i++ {
		slice[i] = i
		arr[i] = i
	}

	updateSlice(slice)
	testArray(arr)
	fmt.Println(slice)
	fmt.Println(arr)

	testArrayPoint(&arr)
	fmt.Println(arr)
}
