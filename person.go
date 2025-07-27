package main

import "fmt"

type person struct {
	name string
	age  int
}

type Employee struct {
	pe         person
	employeeID string
}

func (e Employee) printInfo() {
	fmt.Println(e.employeeID)
	fmt.Println(e.pe)
}

func (e *Employee) updateage() {
	e.pe.age++
}
func testPerson() {
	p := person{name: "sfdaf", age: 20}
	e := Employee{pe: p, employeeID: "123"}
	e.printInfo()

	e.updateage()

	e.printInfo()
}
