package volume

import (
	"testing"
)

const RUNS = 1000000000

func TestDict(t *testing.T) {
	r := Cube{1, 2, 3}
	s := Zylinder{4, 5}

	area_Rec_Wrapper := func(v interface{}) float64 {
		return volume_Cube(v.(Cube))
	}

	area_Sq_Wrapper := func(v interface{}) float64 {
		return volume_Zy(v.(Zylinder))
	}

	x := shape_Value{
		val:    r,
		volume: area_Rec_Wrapper,
	}

	y := shape_Value{
		val:    s,
		volume: area_Sq_Wrapper,
	}

	for i := 0; i < RUNS; i++ {
		_ = sumVolume_Dict(x, y)

		//fmt.Printf("%d \n", z)
	}
}

func TestLookup(t *testing.T) {
	r := Cube{1, 2, 3}
	s := Zylinder{4, 5}

	for i := 0; i < RUNS; i++ {
		_ = sumVolume_Lookup(r, s)

		//fmt.Printf("%d \n", z)
	}
}
