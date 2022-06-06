package mergesortedarray

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	return merge_version2(nums1, m, nums2, n)
}

func merge_version1(nums1 []int, m int, nums2 []int, n int) []int {

	var nums3 []int
	var n1, n2 int
	for k := 0; k < m+n; k++ {
		if n1 < m && n2 < n {
			if nums1[n1] < nums2[n2] {
				nums3 = append(nums3, nums1[n1])
				n1++
				continue
			}
			if nums1[n1] == nums2[n2] {
				nums3 = append(nums3, nums1[n1])
				n1++
				continue
			}
			if nums1[n1] > nums2[n2] {
				nums3 = append(nums3, nums2[n2])
				n2++
				continue
			}
		} else {
			if n1 == m {
				nums3 = append(nums3, nums2[n2])
				n2++
				continue
			}
			if n2 == n {
				nums3 = append(nums3, nums1[n1])
				n1++
				continue
			}
		}

	}
	nums1 = nums3
	return nums1
}

func merge_version2(nums1 []int, m int, nums2 []int, n int) []int {
	i := m - 1
	j := n - 1
	for k := m + n - 1; k >= 0; k-- {
		if i >= 0 && j >= 0 {
			if nums1[i] > nums2[j] {
				nums1[k] = nums1[i]
				i--
			} else if nums1[i] <= nums2[j] {
				nums1[k] = nums2[j]
				j--
			}
		} else if j >= 0 {
			nums1[k] = nums2[j]
			j--
		}
	}
	return nums1
}
