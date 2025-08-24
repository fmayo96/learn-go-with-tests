package structsmethodsinterfaces

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	areaTest := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{2, 4}, 8},
		{Circle{1}, math.Pi},
	}

	for _, tt := range areaTest {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %v, wanted %v", got, tt.want)
		}
	}
}

func TestPerimeter(t *testing.T) {
	areaTest := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{2, 4}, 12},
		{Circle{1}, 2 * math.Pi},
	}

	for _, tt := range areaTest {
		got := tt.shape.Perimeter()
		if got != tt.want {
			t.Errorf("got %v, wanted %v", got, tt.want)
		}
	}
}

