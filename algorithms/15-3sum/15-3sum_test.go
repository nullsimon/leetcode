package threesum

import (
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	input := [][]int{
		{-1, 0, 1, 2, -1, -4},
	}
	want := [][][]int{
		{{-1, -1, 2}, {-1, 0, 1}},
	}
	for i := range input {
		got := threeSum(input[i])
		if !reflect.DeepEqual(got, want[i]) {
			t.Fatalf(`threeSum function failed, want %v, got %v.`, want[i], got)
		}
	}
}
