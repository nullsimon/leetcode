package addtoarrayformofinteger

import (
	"reflect"
	"testing"
)

func TestAddToArrayForm(t *testing.T) {
	input := [][]int{
		{1, 2, 0, 0},
		{2, 7, 4},
		{2, 1, 5},
	}
	k := []int{
		34,
		181,
		806,
	}

	want := [][]int{
		{1, 2, 3, 4},
		{4, 5, 5},
		{1, 0, 2, 1},
	}
	for i := 0; i < len(input); i++ {
		got := addToArrayForm(input[i], k[i])
		if !reflect.DeepEqual(got, want[i]) {
			t.Fatalf(`addToArrayForm function failed, got %v, want %v`, got, want[i])
		}
	}
}
