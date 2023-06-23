package areaTypeAssertion

import (
	"testing"
)

const RUNS = 1000000000

func TestDict(t *testing.T) {
	r := Rectangle{1, 2}
	s := Square{4}

	area_Rec_Wrapper := func(v interface{}) int {
		return area_Rec(v.(Rectangle))
	}

	area_Sq_Wrapper := func(v interface{}) int {
		return area_Sq(v.(Square))
	}

	x := shape_Value{
		val:  r,
		area: area_Rec_Wrapper,
	}

	y := shape_Value{
		val:  s,
		area: area_Sq_Wrapper,
	}

	for i := 0; i < RUNS; i++ {
		_ = sumArea_Dict(x, y)

		//fmt.Printf("%d \n", z)
	}
}

func TestLookup(t *testing.T) {
	r := Rectangle{1, 2}
	s := Square{4}

	for i := 0; i < RUNS; i++ {
		_ = sumArea_Lookup(r, s)

		//fmt.Printf("%d \n", z)
	}
}
