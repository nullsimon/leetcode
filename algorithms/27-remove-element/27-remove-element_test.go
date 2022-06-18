package removeelement

import "testing"

func TestRemoveElement(t *testing.T) {
	input1 := [][]int{
		{3, 2, 2, 3},
		{3},
		{1, 2},
	}
	input2 := []int{
		3,
		3,
		2,
	}
	want := []int{
		2,
		0,
		1,
	}
	for i, _ := range input1 {
		got := removeElement(input1[i], input2[i])
		if got != want[i] {
			t.Fatalf(`removeElement function failed, want %d, got %d`, want[i], got)
		}
	}
}
