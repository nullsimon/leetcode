package topkfrequentelements

import "sort"

func topKFrequent(nums []int, k int) []int {
	var freq = make(map[int]int)

	for _, v := range nums {
		freq[v]++
	}

	sort.Slice(nums, func(i, j int) bool {
		if freq[nums[i]] > freq[nums[j]] {
			return true
		} else if freq[nums[i]] == freq[nums[j]] {
			return nums[i] < nums[j]
		} else if freq[nums[i]] < freq[nums[j]] {
			return false
		}
		return false
	})
	var res []int
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		res = append(res, num)
		if len(res) == k {
			break
		}
		for i+1 < len(nums) && nums[i+1] == num {
			i++
		}
	}

	return res
}
