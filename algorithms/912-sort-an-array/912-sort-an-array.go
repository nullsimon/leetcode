package sortanarray

func sortArray(nums []int) []int {
	//sort(nums, 0, len(nums)-1)
	sortQuick(nums, 0, len(nums)-1)
	return nums
}

func sort(nums []int, low, high int) {
	if low == high {
		return
	}
	mid := low + (high-low)/2
	sort(nums, low, mid)
	sort(nums, mid+1, high)
	merge(nums, low, mid, high)
}

func merge(nums []int, low, mid, high int) {
	var temp = make([]int, high-low+1)

	//two array to one array inverse
	//nums[low:mid]
	//nums[mid+1:high]
	//to temp
	i := low
	j := mid + 1
	for k := 0; k <= high-low; k++ {
		if i == mid+1 {
			temp[k] = nums[j]
			j++
		} else if j == high+1 {
			temp[k] = nums[i]
			i++
		} else if nums[i] <= nums[j] {
			temp[k] = nums[i]
			i++
		} else if nums[j] < nums[i] {
			temp[k] = nums[j]
			j++
		}
	}
	for i := low; i <= high; i++ {
		nums[i] = temp[i-low]
	}

}

func sortQuick(nums []int, low, high int) {
	//shuffle nums
	if low >= high {
		return
	}
	p := partition(nums, low, high)
	sortQuick(nums, low, p-1)
	sortQuick(nums, p+1, high)

}

func partition(nums []int, low, high int) int {
	pivot := nums[low]

	i := low + 1
	j := high
	for i <= j {
		for i < high && nums[i] <= pivot {
			i++
		}
		for j > low && nums[j] > pivot {
			j--
		}
		if i >= j {
			break
		}
		swap(nums, i, j)
	}
	swap(nums, low, j)
	return j
}

func swap(nums []int, i, j int) {
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
}
