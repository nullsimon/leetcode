package removeelement

func removeElement(nums []int, val int) int {
	count := 0
	for _, n := range nums {
		if n != val {
			nums[count] = n
			count++
		}
	}

	return count
}
