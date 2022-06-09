package spiralmatrix

import (
	"reflect"
	"testing"
)

func TestSpiralOrder(t *testing.T) {
	input := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	want := []int{1, 2, 3, 6, 9, 8, 7, 4, 5}
	got := spiralOrder(input)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf(`rotate function failed, got %v, want %v`, got, want)
	}
}
