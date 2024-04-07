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
	got1 := allPathsSourceTarget_BFS(input)

	if !reflect.DeepEqual(got, want) {
		t.Fatalf(`allPathsSourceTarget function failed: want %v, got %v`, want, got)
	}
	if !reflect.DeepEqual(got1, want) {
		t.Fatalf(`allPathsSourceTarget_BFS function failed: want %v, got %v`, want, got1)
	}
}
