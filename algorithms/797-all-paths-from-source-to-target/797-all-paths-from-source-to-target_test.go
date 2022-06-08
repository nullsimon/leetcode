package allpathsfromsourcetotarget

import (
	"reflect"
	"testing"
)

func TestAllPathsSourceTarget(t *testing.T) {
	input := [][]int{
		{1, 2},
		{3},
		{3},
		{},
	}
	want := [][]int{
		{0, 1, 3},
		{0, 2, 3},
	}
	got := allPathsSourceTarget(input)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf(`allPathsSourceTarget function failed: want %v, got %v`, want, got)
	}
}
