package reverseinteger

import "math"

func reverse(x int) int {
	var s []int
	neg := false

	if x == 0 {
		return 0
	}
	if x < 0 {
		neg = true
		x = x / -1
	}
	for x > 0 {
		i := x % 10
		x = x / 10
		s = append(s, i)
	}

	var reverse = 0
	for _, j := range s {
		reverse = reverse*10 + j
		if reverse > math.MaxInt32 {
			return 0
		}
		//    res := math.Pow(2, 10)
	}
	if neg {
		reverse = -reverse
		if reverse < math.MinInt32 {
			return 0
		}
	}
	return reverse
}
