package maximumsubarray

import "math"

// sliding window solution
func maxSubArray(nums []int) int {
	var sum int
	var res = math.MinInt64
	left, right := 0, 0

	for right < len(nums) {
		// add new element
		r := nums[right]
		sum = sum + r
		right++
		if sum > res {
			res = sum
		}
		// shrink, why < 0 ,because if all of them are negative, so we need find the smallest
		for sum < 0 {
			sum -= nums[left]
			left++
		}
	}
	return res
}

// KadanesAlgorithm
func maxSubArray1(nums []int) int {
	sum, max := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		// we divide all nums to three part
		// left_big_sum  num[i]  right_part
		// so choose left or num[i] to continue bigest?
		sum = Max(sum+nums[i], nums[i])
		max = Max(sum, max)
	}
	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
