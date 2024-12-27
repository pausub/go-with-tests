package structs

import "math"

type Shape interface {
	Area() float64
}

// no neeed to define "implements Shape"
// interface resolution is implicit
type Rectangle struct {
	X, Y float64
}

type Circle struct {
	R float64
}

type Triangle struct {
	Base , Height float64
}

func Perimeter(r Rectangle) float64 {
	return 2 * (r.X + r.Y)
}

// syntax: func (receiverName: receiverType) MethodName(args)
// receiver convention: first lovercase letter of type
func (r Rectangle) Area() float64 {
	return r.X * r.Y
}

func (c Circle) Area() float64 {
	return math.Pow(c.R, 2) * math.Pi
}

func (t Triangle) Area() float64 {
	return t.Base * t.Height / 2
}
