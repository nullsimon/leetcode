package medianoftwosortedarrays

// O(m+n) version
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var nums3 []int
	i, j := 0, 0
	for i < len(nums1) || j < len(nums2) {
		if i < len(nums1) && j < len(nums2) {
			if nums1[i] <= nums2[j] {
				nums3 = append(nums3, nums1[i])
				i++
				continue
			}
			if nums1[i] > nums2[j] {
				nums3 = append(nums3, nums2[j])
				j++
				continue
			}
		}
		if i < len(nums1) {
			nums3 = append(nums3, nums1[i])
			i++
			continue

		}
		if j < len(nums2) {
			nums3 = append(nums3, nums2[j])
			j++
			continue
		}
	}
	if (i+j)%2 == 0 {
		middle1 := (i+j)/2 - 1
		middle2 := (i + j) / 2
		return float64((nums3[middle1] + nums3[middle2])) / 2
	} else {
		middle := (i + j) / 2
		return float64(nums3[middle])
	}
}
