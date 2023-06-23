package areaTypeAssertion

import "fmt"

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

// Dictionary translation

type shape_Value struct {
	val  interface{}
	area func(interface{}) int
}

func sumArea_Dict(x, y shape_Value) int {
	xVal, ok := x.val.(shape)
	if !ok {
		fmt.Println("type of x.val not accepted")
		return 0
	}
	yVal, ok := y.val.(shape)
	if !ok {
		fmt.Println("type of y.val not accepted")
		return 0
	}

	return x.area(xVal) + y.area(yVal)
}
