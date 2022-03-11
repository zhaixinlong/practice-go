package sort_slice

// 选择排序
func SelectionSort(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		minIdx := i
		for j := i + 1; j < len(nums); j++ {
			if nums[minIdx] > nums[j] {
				minIdx = j
			}
		}

		if minIdx == i {
			continue
		}

		minValue := nums[minIdx]
		nums[minIdx] = nums[i]
		nums[i] = minValue
	}
	return nums
}
