package structs

import "math"

// Shape represents an 2 or 3 dimensinal object
type Shape interface {
	Area() float64
}

// Rectangle represents a 4 sided object with equal width and height
type Rectangle struct {
	Width  float64
	Height float64
}

// Perimeter returns perimiter of rectangle of a given width and height
func Perimeter(r Rectangle) float64 {
	return 2*r.Width + 2*r.Height
}

// Area returns the area of a rectangle of a given width and height
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle represents a circle shape with a given radius
type Circle struct {
	Radius float64
}

// Area returns the area of a circle of a given radius
func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

// Triangle represents a right triangle with a given base and height
type Triangle struct {
	Base   float64
	Height float64
}

// Area returns the area of a right triangle with a given base and height
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}
