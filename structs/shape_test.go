package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f but wanted %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	areaTest := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 10.0, Height: 10.0}, want: 100.0},
		{name: "Circle", shape: Circle{Radius: 10}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, want: 36},
	}

	for _, tt := range areaTest {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("%#v got %.2f but want %.2f", tt.shape, got, tt.want)
		}

	}
}
