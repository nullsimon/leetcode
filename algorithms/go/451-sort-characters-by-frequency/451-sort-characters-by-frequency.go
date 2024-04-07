package sort

import (
	"sort"
)

//copy from https://leetcode.com/problems/sort-characters-by-frequency/discuss/2062345/Go.-Golang.-Short-simple-sorting.
func frequencySort(s string) string {
	//since string are uint8 byte slice
	b := []byte(s)
	frequency := make([]int, 128)
	for _, ch := range b {
		frequency[ch]++
	}

	//sort slice ,using internal package
	sort.Slice(b, func(i, j int) bool {
		freqI, freqJ := frequency[b[i]], frequency[b[j]]
		if freqI != freqJ {
			return freqI > freqJ
		}
		return b[i] > b[j]
	})

	return string(b)
}

// ordinary version
func frequencySort1(s string) string {
	var count = make(map[string]int)
	for _, v := range []byte(s) {
		vs := string(v)
		if _, ok := count[vs]; ok {
			count[vs] = count[vs] + 1
		} else {
			count[vs] = 1
		}
	}
	res := ""
	ranked := rankMapStringInt(count)

	for i := 0; i < len(ranked); i++ {
		c := count[ranked[i]]
		for j := 0; j < c; j++ {
			res += ranked[i]
		}
	}
	return res

}

func rankMapStringInt(values map[string]int) []string {
	type kv struct {
		Key   string
		Value int
	}
	var ss []kv
	for k, v := range values {
		ss = append(ss, kv{k, v})
	}
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})
	ranked := make([]string, len(values))
	for i, kv := range ss {
		ranked[i] = kv.Key
	}
	return ranked
}
