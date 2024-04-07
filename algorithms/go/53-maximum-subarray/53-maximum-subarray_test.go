package maximumsubarray

import "testing"

func TestMaxSubArray(t *testing.T) {
	input := [][]int{
		{-2, 1, -3, 4, -1, 2, 1, -5, 4},
		{1},
		{5, 4, -1, 7, 8},
	}
	want := []int{
		6,
		1,
		23,
	}
	for i := range input {
		got := maxSubArray(input[i])
		got1 := maxSubArray1(input[i])
		if got != want[i] || got1 != want[i] {
			t.Fatalf(`maxSubArray function failed, want %d, got %d, got1 %d`, want[i], got, got1)
		}
	}
}
