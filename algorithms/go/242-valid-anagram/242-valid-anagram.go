package validanagram

import "reflect"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var window, need = make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(s); i++ {
		window[s[i]]++
		need[t[i]]++
	}
	if !reflect.DeepEqual(window, need) {
		return false
	}
	return true
}
