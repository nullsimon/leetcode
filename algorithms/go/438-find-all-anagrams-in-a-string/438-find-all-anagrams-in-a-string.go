package findallanagramsinastring

import (
	"reflect"
)

func findAnagrams(s string, p string) []int {
	var window = make(map[byte]int)
	var need = make(map[byte]int)
	for k := range p {
		need[p[k]]++
	}
	var pl = len(p)

	left, right := 0, 0
	var res []int

	for right < len(s) {
		c := s[right]
		right++
		window[c]++

		for right-left > pl {
			l := s[left]
			left++
			window[l]--
			if window[l] == 0 {
				delete(window, l)
			}
		}
		if right-left == pl {
			if reflect.DeepEqual(window, need) {
				res = append(res, left)
			}
		}

	}
	return res
}
