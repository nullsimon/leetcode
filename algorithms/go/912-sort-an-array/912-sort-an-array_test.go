package sortanarray

import (
	"reflect"
	"testing"
)

func TestSortArray(t *testing.T) {
	input := [][]int{
		{5, 2, 3, 1},
		{5, 1, 1, 2, 0, 0},
		{0, 0},
	}
	want := [][]int{
		{1, 2, 3, 5},
		{0, 0, 1, 1, 2, 5},
		{0, 0},
	}

	for i, v := range input {
		got := sortArray(v)
		if !reflect.DeepEqual(got, want[i]) {
			t.Fatalf(`sortArray function failed, want %v, got %v .`, want[i], got)
		}
	}

}
