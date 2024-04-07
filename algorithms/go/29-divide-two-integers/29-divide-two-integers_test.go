package dividetwointegers

import "testing"

func TestDivide(t *testing.T) {
	input := [][]int{
		{10, 3},
		{7, -3},
		{-2147483648, -1},
	}
	want := []int{
		3,
		-2,
		2147483647,
	}
	for i := range input {
		got := divide(input[i][0], input[i][1])
		if got != want[i] {
			t.Fatalf(`divide function failed, want %d, got %d`, want[i], got)
		}
	}
}
