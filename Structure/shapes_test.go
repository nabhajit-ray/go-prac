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
	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g and want %g", got, want)

		}
	}
	t.Run("Testing rectangle area", func(t *testing.T) {
		rectangle := Rectangle{20.0, 40.0}
		checkArea(t, rectangle, 800.0)
	})
	t.Run("Testing circle area", func(t *testing.T) {
		circle := Circle{20.0}
		checkArea(t, circle, 1256.6370614359173)

	})
}
