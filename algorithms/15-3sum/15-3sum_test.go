package threesum

import (
	"reflect"
	"sort"
	"testing"
)

func TestThreeSum(t *testing.T) {
	input := [][]int{
		{-1, 0, 1, 2, -1, -4},
	}
	want := [][][]int{
		{{-1, 0, 1}, {-1, -1, 2}},
	}
	for i := range input {
		got := threeSum2(input[i])
		for k := range want[i] {
			sort.Ints(want[i][k])
			sort.Ints(got[k])
		}
		if !reflect.DeepEqual(got, want[i]) {
			t.Fatalf(`threeSum function failed, want %v, got %v.`, want[i], got)
		}
	}
}
