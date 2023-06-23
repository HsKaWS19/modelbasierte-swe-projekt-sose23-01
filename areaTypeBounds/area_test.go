package areaTypeBounds

import (
	"testing"
)

const RUNS = 1000000000

func TestDict(t *testing.T) {
	r := Rectangle{1, 2}
	s := Square{4}

	/*
		area_Rec_Wrapper := func(v Rectangle) int {
			return area_Rec(v.(Rectangle))
		}

		area_Sq_Wrapper := func(v interface{}) int {
			return area_Sq(v.(Square))
		}
	*/

	x := shape_Value[Rectangle]{
		val: r,
		area: func(v Rectangle) int {
			return area_Rec(v)
		},
	}

	y := shape_Value[Square]{
		val: s,
		area: func(v Square) int {
			return area_Sq(v)
		},
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
