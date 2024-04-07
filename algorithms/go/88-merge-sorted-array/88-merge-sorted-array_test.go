package mergesortedarray

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	want := []int{1, 2, 2, 3, 5, 6}
	got := merge(nums1, m, nums2, n)
	if !reflect.DeepEqual(want, got) {
		t.Fatalf(`merge function not working as expected, want %v, got %v`, want, got)
	}
}
