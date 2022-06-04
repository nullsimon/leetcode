package mergesortedarray

import "fmt"

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
	fmt.Println(nums3)
	nums1 = nums3
	fmt.Println(nums1)
	return nums1
}

func merge_version2(nums1 []int, m int, nums2 []int, n int) []int {
	var n1, n2 int
	for k := 0; k < m+n; k++ {
		if n1 < m && n2 < n {
			if nums1[n1] <= nums2[n2] {
				n1++
				continue
			}
			if nums1[n1] > nums2[n2] {
				nums3 := nums1[n1 : m+n-1]
				nums1 = append(nums1[:n1], nums2[n2])
				nums1 = append(nums1, nums3...)
				n2++
				continue
			}
		}
	}
	return nums1
}
