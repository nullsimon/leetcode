package spiralmatrixii

import (
	"reflect"
	"testing"
)

func TestGenerateMatrix(t *testing.T) {
	input := 3
	want := [][]int{
		{1, 2, 3},
		{8, 9, 4},
		{7, 6, 5},
	}
	got := generateMatrix(input)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf(`generateMatrix function failed，want %v, got %v.`, want, got)
	}
}
