package lettercombinationsofaphonenumber

import (
	"strconv"
	"strings"
)

func letterCombinations(digits string) []string {
	var ls = [][]string{
		{"a", "b", "c"},
		{"d", "e", "f"},
		{"g", "h", "i"},
		{"j", "k", "l"},
		{"m", "n", "o"},
		{"p", "q", "r", "s"},
		{"t", "u", "v"},
		{"w", "x", "y", "z"},
	}

	list := strings.Split(digits, "")
	if len(list) == 0 {
		return nil
	}
	// one build first
	i1, _ := strconv.Atoi(list[0])
	l1 := ls[i1-2]
	if len(list) == 1 {
		return l1
	}
	// two
	i2, _ := strconv.Atoi(list[1])
	l2 := ls[i2-2]
	var res []string
	for _, m := range l1 {
		for _, n := range l2 {
			res = append(res, m+n)
		}
	}

	if len(list) == 2 {
		return res
	}

	// three = two + one
	i3, _ := strconv.Atoi(list[2])
	l3 := ls[i3-2]
	var res3 []string
	for _, m := range res {
		for _, n := range l3 {
			res3 = append(res3, m+n)
		}
	}

	if len(list) == 3 {
		return res3
	}
	// four = three + one
	i4, _ := strconv.Atoi(list[3])
	l4 := ls[i4-2]
	var res4 []string
	for _, m := range res3 {
		for _, n := range l4 {
			res4 = append(res4, m+n)
		}
	}
	return res4
}
