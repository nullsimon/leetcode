package rotateimage

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	input := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	want := [][]int{
		{7, 4, 1},
		{8, 5, 2},
		{9, 6, 3},
	}
	got := rotate(input)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf(`rotate function failed, got %v, want %v`, got, want)
	}
}
