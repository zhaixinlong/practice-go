package sort_slice

import "fmt"

// 冒泡排序
func BubbleSort(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		fmt.Printf("temp i %d \n", nums[i])
		// 设定一个标记，若为true，则表示此次循环没有进行交换，也就是待排序列已经有序，排序已经完成。
		flag := true

		for j := i + 1; j < len(nums); j++ {
			temp := nums[i]
			fmt.Printf("i %d j %d temp %d target %d \n", i, j, temp, nums[j])
			if temp > nums[j] {
				nums[i] = nums[j]
				nums[j] = temp

				flag = false
			}
		}
		fmt.Printf("------")
		fmt.Printf("nums %v len %d \n", nums, len(nums))

		if flag {
			break
		}
	}
	return nums
}
