package removeduplicatesfromsortedarray

func removeDuplicates(nums []int) int {
	k := 0
	for i := 0; i < len(nums); i++ {
		nums[k] = nums[i]
		for i < len(nums)-1 && nums[k] == nums[i+1] {
			i++
		}
		k++
	}
	return k
}
