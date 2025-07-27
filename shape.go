package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	perimeter() float64
}

type circle struct {
	radius float64
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c *circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

type square struct {
	height float64
	width  float64
}

func (s *square) area() float64 {
	return s.height * s.width
}

func (s *square) perimeter() float64 {
	return (s.height + s.width) * 2
}

func testShape() {
	var s shape
	c := circle{radius: 10}

	// s = c // ❌ 编译错误，circle（值）未实现接口（需要指针）
	s = &c // ✅ 正确

	var s2 shape
	sq := square{height: 5, width: 5}
	s2 = &sq

	fmt.Println(s.area())
	fmt.Println(s2.perimeter())
}
