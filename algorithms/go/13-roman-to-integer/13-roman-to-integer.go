package romanToInteger

import "fmt"

func romanToInt(s string) int {
	var romansInt = map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	res := 0
	var pre string
	for _, v := range s {
		v := string(v)
		fmt.Println(v)
		if pre != "" {
			if pre == v {
				res += romansInt[v]
			}
			if pre != v {
				if romansInt[pre] < romansInt[v] {
					res -= romansInt[pre] * 2
					res += romansInt[v]
				}
				if romansInt[pre] == romansInt[v] {
					res += romansInt[v]
				}
				if romansInt[pre] > romansInt[v] {
					res += romansInt[v]
				}
			}
		}
		if pre == "" {
			res += romansInt[v]
		}
		pre = v
	}

	return res
}
