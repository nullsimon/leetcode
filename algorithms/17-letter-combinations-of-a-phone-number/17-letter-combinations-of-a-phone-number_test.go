package lettercombinationsofaphonenumber

import (
	"reflect"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	input := []string{
		"23",
		"",
		"2",
	}
	want := [][]string{
		{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		nil,
		{"a", "b", "c"},
	}
	for i := range input {
		got := letterCombinations(input[i])
		if !reflect.DeepEqual(got, want[i]) {
			t.Fatalf("function letterCombinations failed: want %v, got %v", want[i], got)
		}
	}

}
