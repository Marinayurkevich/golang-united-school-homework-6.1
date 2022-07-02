package golang_united_school_homework

// Rectangle must satisfy to Shape interface
type Rectangle struct {
	Height, Weight float64
}

func (r Rectangle) CalcArea() float64 {
	s := r.Height * r.Weight
	return s
}

func (r Rectangle) CalcPerimeter() float64 {
	p := 2 * (r.Height + r.Weight)
	return p
}
