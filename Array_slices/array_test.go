package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Adding all the elements in an slice", func(t *testing.T) {

		numbers := []int{1, 2, 3, 4, 5, 6}
		sum := SumArray(numbers)
		expected := 21
		if sum != expected {
			t.Errorf("sum %d expected %d", sum, expected)
		}
	})

}
func TestSumAll(t *testing.T) {

	t.Run("Adding up all slices", func(t *testing.T) {
		slice1 := []int{1, 3, 4, 2}
		slice2 := []int{1, 5, 8}
		sum := SumAll(slice1, slice2)
		expected := []int{10, 14}
		if !reflect.DeepEqual(sum, expected) {
			t.Errorf("sum %v expected %v", sum, expected)
		}

	})
}
