package sqrtx

// simple O(n) version
func mySqrt1(x int) int {
	if x == 0 {
		return 0
	}
	i := 1
	for i <= x {
		if i*i == x {
			return i
		} else if i*i < x {
			i++
		} else if i*i > x {
			i--
			return i
		}
	}
	return i
}

// binary search version O(log n)
func mySqrt(x int) int {
	// base cases
	if x <= 1 {
		return x
	}
	low, high := 0, x
	var mid int
	for low < high {
		mid = low + (high-low)/2
		if mid == low {
			return mid
		}
		if mid*mid == x {
			return mid
		} else if mid*mid < x {
			low = mid
		} else if mid*mid > x {
			high = mid
		}
	}
	return low
}
