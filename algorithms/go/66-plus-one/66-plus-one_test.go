package plusone

import (
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {
	input := [][]int{
		{1, 2, 3},
		{9},
	}

	want := [][]int{
		{1, 2, 4},
		{1, 0},
	}

	for i := 0; i < len(input); i++ {
		got := plusOne(input[i])
		if !reflect.DeepEqual(got, want[i]) {
			t.Errorf("got: %v, want: %v", got, want[i])
		}
	}
}
