package findallanagramsinastring

import (
	"reflect"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	input := []string{
		"cbaebabacd",
		"abab",
		"baa",
	}
	input2 := []string{
		"abc",
		"ab",
		"aa",
	}

	want := [][]int{
		{0, 6},
		{0, 1, 2},
		{1},
	}
	for i := range input {
		got := findAnagrams(input[i], input2[i])
		if !reflect.DeepEqual(got, want[i]) {
			t.Fatalf(`findAnagrams function failed, want %v, got %v.`, want[i], got)
		}
	}

}
