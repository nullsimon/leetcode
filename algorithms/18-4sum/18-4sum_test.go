package foursum

import (
	"reflect"
	"sort"
	"testing"
)

func TestFourSum(t *testing.T) {
	input := [][]int{
		{1, 0, -1, 0, -2, 2},
		{2, 2, 2, 2},
	}
	target := []int{
		0,
		8,
	}
	want := [][][]int{
		{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}},

		{{2, 2, 2, 2}},
	}
	for i := range input {
		got := fourSum(input[i], target[i])
		for k := range want[i] {
			sort.Ints(want[i][k])
			sort.Ints(got[k])
		}
		if !reflect.DeepEqual(got, want[i]) {
			t.Fatalf(`fourSum function failed, want %v, got %v.`, want[i], got)
		}
	}
}
