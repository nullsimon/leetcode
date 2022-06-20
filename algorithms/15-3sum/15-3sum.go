package threesum

import "sort"

func threeSum(nums []int) [][]int {
	var res [][]int
	var hisMap = make(map[[3]int]bool)
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		var1 := nums[i]
		for j := i + 1; j < len(nums)-1; j++ {
			var2 := nums[j]
			for k := j + 1; k < len(nums); k++ {
				var3 := nums[k]
				if var1+var2+var3 == 0 {
					values := [3]int{var1, var2, var3}

					if _, ok := hisMap[values]; !ok {
						var s []int
						s = values[0:3]
						res = append(res, s)
						hisMap[values] = true

					}
				}
			}
		}
	}
	return res
}

func threeSum1(nums []int) [][]int {
	var res [][]int
	var hisMap = make(map[[3]int]bool)
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		target := 0 - nums[i]
		res1 := twoSum(nums[i+1:], target)
		for j := range res1 {
			res1[j] = append(res1[j], nums[i])
			sort.Ints(res1[j])
			var values [3]int
			for k := 0; k < 3; k++ {
				values[k] = res1[j][k]
			}
			if _, ok := hisMap[values]; !ok {
				res = append(res, res1[j])
				hisMap[values] = true
			}
		}

	}
	return res

}

func threeSum2(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		target := 0 - nums[i]
		res1 := twoSum(nums[i+1:], target)
		for j := range res1 {
			res1[j] = append(res1[j], nums[i])
			res = append(res, res1[j])
		}
		// remove dup
		for i < len(nums)-1 && nums[i] == nums[i+1] {
			i++
		}

	}
	return res

}

func twoSum(nums []int, target int) [][]int {

	var m = make(map[int]int)
	var res [][]int

	for _, num := range nums {
		if _, ok := m[num]; ok {
			res1 := []int{num, m[num]}
			res = append(res, res1)
		}
		m[target-num] = num
	}

	return res
}
