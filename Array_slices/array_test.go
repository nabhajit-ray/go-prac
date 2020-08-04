package main

import "testing"

func TestSum(t *testing.T) {
	t.Run("Adding all the elements in an array", func(t *testing.T) {

		numbers := []int{1, 2, 3, 4, 5, 6}
		sum := SumArray(numbers)
		expected := 21
		if sum != expected {
			t.Errorf("sum %d expected %d", sum, expected)
		}
	})

}
