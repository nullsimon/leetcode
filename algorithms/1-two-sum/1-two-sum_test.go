package twosum

import (
	"reflect"
	"sort"
	"testing"
)

func TestTwoSum(t *testing.T) {
	input := [][]int{
		{2, 7, 11, 15},
		{3, 2, 4},
		{3, 3},
	}
	target := []int{
		9,
		6,
		6,
	}
	want := [][]int{
		{0, 1},
		{1, 2},
		{0, 1},
	}
	for i := range input {
		got := twoSum1(input[i], target[i])
		sort.Ints(got)
		if !reflect.DeepEqual(got, want[i]) {
			t.Fatalf(`twoSum function failed,input %v, target %v,want %v, got %v`, input[i], target[i], want[i], got)
		}
	}
}
