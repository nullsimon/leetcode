package twosum

func twoSum(nums []int, target int) []int {
	for i, _ := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil

}

func twoSum1(nums []int, target int) []int {

	var m = make(map[int]int)

	for i, num := range nums {
		if val, ok := m[num]; ok {
			return []int{i, val}
		}
		m[target-num] = i
	}
	return nil
}
