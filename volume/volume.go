package volume

import "math"

type Cube struct {
	length float64
	width  float64
	high   float64
}

type Zylinder struct {
	radius float64
	high   float64
}

func (r Cube) volume() float64 {
	return r.length * r.width * r.high
}

func (s Zylinder) volume() float64 {
	return s.high * s.radius * s.radius * math.Pi
}

type shape interface {
	volume() int
}

// Introducing unique function names for overloaded methods

func volume_Cube(r Cube) float64 {
	return r.length * r.width * r.high
}

func volume_Zy(s Zylinder) float64 {
	return s.high * s.radius * s.radius * math.Pi
}

// Run-time method lookup

func volume_Lookup(x interface{}) float64 {
	var y float64

	switch v := x.(type) {
	case Zylinder:
		y = volume_Zy(v)
	case Cube:
		y = volume_Cube(v)
	}
	return y

}

func sumVolume_Lookup(x, y interface{}) float64 {
	return volume_Lookup(x) + volume_Lookup(y)
}

// Dictionary translation

type shape_Value struct {
	val    interface{}
	volume func(interface{}) float64
}

func sumVolume_Dict(x, y shape_Value) float64 {
	return x.volume(x.val) + y.volume(y.val)
}
