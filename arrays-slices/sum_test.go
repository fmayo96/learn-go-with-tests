package arraysslices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Sum slice of 5 numbers", func(t *testing.T) {
		arr := []int{1, 2, 3, 4, 5}
		got := Sum(arr)
		want := 15
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("Sum slice of 3 numbers", func(t *testing.T) {
		arr := []int{1, 2, 3}
		got := Sum(arr)
		want := 6
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{3, 6}, []int{1, 1})
	want := []int{9, 2}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
