package foursum

import "sort"

func fourSum(nums []int, target int) [][]int {
	var res [][]int
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		target2 := target - nums[i]
		res1 := threeSum2(nums[i+1:], target2)
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

func threeSum2(nums []int, target int) [][]int {
	var res [][]int
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		target1 := target - nums[i]
		res1 := twoSum(nums[i+1:], target1)
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

	sort.Ints(nums)
	var res [][]int
	var left, right = 0, len(nums) - 1
	for left < right {
		var sum = nums[left] + nums[right]
		l := nums[left]
		r := nums[right]
		if sum > target {
			for left < right && nums[right] == r {
				right--
			}
		} else if sum < target {
			for left < right && nums[left] == l {
				left++
			}
		} else if sum == target {
			res = append(res, []int{nums[left], nums[right]})
			for left < right && nums[right] == r {
				right--
			}
			for left < right && nums[left] == l {
				left++
			}
		}
	}
	return res
}
