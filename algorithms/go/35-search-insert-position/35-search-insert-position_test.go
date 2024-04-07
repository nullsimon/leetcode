package searchinsertposition

import "testing"

func TestSearchInsert(t *testing.T) {
	input := [][]int{
		{1, 3, 5, 6},
		{1, 3, 5, 6},
		{1, 3, 5, 6},
	}
	target := []int{
		5,
		2,
		7,
	}
	want := []int{
		2,
		1,
		4,
	}
	for i := range input {
		got := searchInsert(input[i], target[i])
		if got != want[i] {
			t.Fatalf(`searchInsert function failed,input %v, target %d, want %d, got %d.`, input[i], target[i], want[i], got)
		}
	}
}
