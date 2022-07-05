package validperfectsquare

func isPerfectSquare(num int) bool {
	return mySqrt(num)
}

// binary search version O(log n)
func mySqrt(x int) bool {
	// base cases
	if x <= 1 {
		return true
	}
	low, high := 0, x
	var mid int
	for low < high {
		mid = low + (high-low)/2
		if mid == low {
			return false
		}
		if mid*mid == x {
			return true
		} else if mid*mid < x {
			low = mid
		} else if mid*mid > x {
			high = mid
		}
	}
	return false
}
