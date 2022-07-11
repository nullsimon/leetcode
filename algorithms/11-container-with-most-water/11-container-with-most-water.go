package containerwithmostwater

func maxArea(height []int) int {
	// min(height[left], height[right]) * (right - left)
	var max int
	var amount int
	for i := 0; i < len(height)-1; i++ {
		for j := 1; j < len(height); j++ {
			amount = (j - i) * min(height[i], height[j])
			if amount > max {
				max = amount
			}
		}
	}
	return max
}

func maxArea1(height []int) int {
	var max int
	left, right := 0, len(height)-1
	for left < right {
		if height[left] < height[right] {
			max = maxInt(max, (right-left)*height[left])
			left++
		} else {
			max = maxInt(max, (right-left)*height[right])
			right--
		}
	}
	return max
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
