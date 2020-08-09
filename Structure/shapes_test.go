package main

import "testing"

func TestPerimeter(t *testing.T) {

	t.Run("Testing rectangle perimeter", func(t *testing.T) {
		rectangle := Rectangle{10.0, 20.0}
		got := Perimeter(rectangle)
		want := 60.0
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}

	})
}
func TestArea(t *testing.T) {

	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{20}, 1256.6370614359173},
		{Triangle{12, 6}, 36.0},
	}
	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}

}
