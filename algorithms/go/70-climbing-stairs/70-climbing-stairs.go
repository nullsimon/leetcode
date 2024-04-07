package climbingstairs

func climbStairs(n int) int {
	var m = make(map[int]int)
	return dp(n, m)

}

func dp(n int, m map[int]int) int {
	// base case
	if n <= 2 {
		return n
	}
	if _, ok := m[n]; ok {
		return m[n]
	}
	m[n] = dp(n-1, m) + dp(n-2, m)
	return m[n]
}
