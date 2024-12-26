package structs

import "math"

type Rectangle struct {
	X, Y float64
}

type Circle struct {
	R float64
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
	return c.R * c.R * math.Pi
}
