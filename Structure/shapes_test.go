package main

import "testing"

func TestPerimeter(t *testing.T) {

	t.Run("Testing rectangle shape", func(t *testing.T) {
		got := Perimeter(10.0, 20.0)
		want := 60.0
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}

	})
}
func TestArea(t *testing.T) {
	got := Area(10.0, 30.0)
	want := 300.0
	if got != want {
		t.Errorf("got %.2f and want %.2f", got, want)
	}
}
