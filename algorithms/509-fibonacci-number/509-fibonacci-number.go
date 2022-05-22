package fibonaccinumber

func fib1(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

//fib with dp memo table, using map to store result
func fib(n int) int {
	var memo = make(map[int]int, n)
	return helper(memo, n)
}

func helper(memo map[int]int, n int) int {
	if n == 0 || n == 1 {
		return n
	}
	if memo[n] != 0 {
		return memo[n]
	}
	memo[n] = helper(memo, n-1) + helper(memo, n-2)
	return memo[n]
}
