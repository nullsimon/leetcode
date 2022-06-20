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
