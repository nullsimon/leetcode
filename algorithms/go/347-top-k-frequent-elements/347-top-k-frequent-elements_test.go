package topkfrequentelements

import (
	"reflect"
	"sort"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	input := [][]int{
		{3, 0, 1, 0},
		{1, 1, 1, 2, 2, 3},
		{3, 2, 3, 1, 2, 4, 5, 5, 6, 7, 7, 8, 2, 3, 1, 1, 1, 10, 11, 5, 6, 2, 4, 7, 8, 5, 6},
	}
	ks := []int{
		1,
		2,
		10,
	}
	want := [][]int{
		{0},
		{1, 2},
		{1, 2, 5, 3, 6, 7, 4, 8, 10, 11},
	}
	for i := range input {
		got := topKFrequent(input[i], ks[i])
		sort.Ints(got)
		sort.Ints(want[i])
		if !reflect.DeepEqual(got, want[i]) {
			t.Fatalf(`topKFrequent function failed, got %v, want %v`, got, want[i])
		}
	}
}
