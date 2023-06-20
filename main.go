package main

import "fmt"

// Example

type Rectangle struct {
	length int
	width  int
}

type Square struct {
	length int
}

func (r Rectangle) area() int {
	return r.length * r.width
}

func (s Square) area() int {
	return s.length * s.length
}

type shape interface {
	area() int
}

func sumArea(x, y shape) int {
	return x.area() + y.area()
}

// Introducing unique function names for overloaded methods

func area_Rec(r Rectangle) int {
	return r.length * r.width
}

func area_Sq(s Square) int {
	return s.length * s.length
}

// Run-time method lookup

func area_Lookup(x interface{}) int {
	var y int

	switch v := x.(type) {
	case Square:
		y = area_Sq(v)
	case Rectangle:
		y = area_Rec(v)
	}
	return y

}

func sumArea_Lookup(x, y interface{}) int {
	return area_Lookup(x) + area_Lookup(y)
}

func test_Lookup() {
	var r Rectangle = Rectangle{1, 2}
	var s Square = Square{3}

	x1 := area_Rec(r) + area_Sq(s)

	fmt.Printf("%d \n", x1)

	x2 := sumArea_Lookup(r, s)
	// Rectangle <= interface{}
	// Square <= interface{}

	fmt.Printf("%d \n", x2)
}

// Dictionary translation

type shape_Value struct {
	val  interface{}
	area func(interface{}) int
}

func sumArea_Dict(x, y shape_Value) int {
	return x.area(x.val) + y.area(y.val)
}

func main() {

	//test_Lookup()

	//test_Dict()

}
