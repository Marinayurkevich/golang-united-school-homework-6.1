package golang_united_school_homework

import "math"

// Triangle must satisfy to Shape interface
type Triangle struct {
	Side float64
}

func (t Triangle) CalcArea() float64 {
	s := (math.Pow(t.Side, 2) * math.Sqrt(3)) / 4
	return s
}

func (t Triangle) CalcPerimeter() float64 {
	p := 3 * t.Side
	return p
}
