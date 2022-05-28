package coinchange

import "math"

func coinChange(coins []int, amount int) int {

	var dpMaps = make(map[int]int)
	return dp(coins, amount, dpMaps)
}

func dp(coins []int, n int, dpMaps map[int]int) int {

	if n == 0 {
		return 0
	}
	if n < 0 {
		return -1
	}

	if dpMaps[n] > 0 {
		return dpMaps[n]
	}

	res := math.MaxInt32 - 1

	for _, coin := range coins {
		sub := dp(coins, n-coin, dpMaps)
		if sub < 0 {
			continue
		}
		res = min(res, sub+1)
	}
	if res == math.MaxInt32-1 {
		return -1
	}
	dpMaps[n] = res
	return res

}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
