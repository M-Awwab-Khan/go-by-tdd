package shapes

import "math";

type Rectangle struct {
	Width float64
	Height float64
}

type Shape interface {
	Area() float64
}

func (rect Rectangle) Perimeter() float64 {
	return 2 * (rect.Height + rect.Width);
}

func (rect Rectangle) Area() float64 {
	return rect.Width * rect.Height;
}

type Circle struct {
	Radius float64
}


func (c Circle) Perimeter() float64 {
	return 2 * math.Pi  * c.Radius;
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2);
}

type Triangle struct {
	Base float64
	Height float64
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height;
}