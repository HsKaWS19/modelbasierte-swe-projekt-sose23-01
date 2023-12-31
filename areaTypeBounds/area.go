package areaTypeBounds

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

func area_Lookup[T shape](x T) int {
	return x.area()
}

func sumArea_Lookup[T, R shape](x T, y R) int {
	return area_Lookup(x) + area_Lookup(y)
}

// Dictionary translation

type shape_Value[T shape] struct {
	val  T
	area func(T) int
}

func sumArea_Dict[T, R shape](x shape_Value[T], y shape_Value[R]) int {
	return x.area(x.val) + y.area(y.val)
}
