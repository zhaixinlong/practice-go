package sort_slice

// 插入排序
func InsertionSort(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		tmp := nums[i]

		for j := i; j > 0; j-- {
			if tmp > nums[j-1] {
				swap(nums, j, j-1)
			} else {
				break
			}
		}
	}
	return nums
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
