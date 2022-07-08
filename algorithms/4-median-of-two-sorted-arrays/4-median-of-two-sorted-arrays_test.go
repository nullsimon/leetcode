package medianoftwosortedarrays

import "testing"

func TestFindMedianSortedArrays(t *testing.T) {
	input := [][]int{
		{1, 3},
		{1, 2},
		{1, 2, 3, 4, 5},
	}
	input2 := [][]int{
		{2},
		{3, 4},
		{1, 2, 3, 4, 5},
	}
	want := []float64{
		2.0,
		2.5,
		3.0,
	}
	for i, v := range input {
		if got := findMedianSortedArrays(v, input2[i]); got != want[i] {
			t.Errorf("findMedianSortedArrays(%v, %v) = %v; want %v", v, input2[i], got, want[i])
		}
	}
}
