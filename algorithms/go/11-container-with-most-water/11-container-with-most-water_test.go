package containerwithmostwater

import "testing"

func TestMaxArea(t *testing.T) {
	input := [][]int{
		{1, 8, 6, 2, 5, 4, 8, 3, 7},
		{1, 1},
	}
	want := []int{
		49,
		1,
	}
	for i, v := range input {
		if got := maxArea1(v); got != want[i] {
			t.Errorf("maxArea(%v) = %v; want %v", v, got, want[i])
		}
	}
}
