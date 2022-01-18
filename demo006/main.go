package main

import (
	"fmt"
)

// 排序

// 冒泡排序
func sortBubble(nums []int) []int {
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

// 选择排序
func sortSelection(nums []int) []int {
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

// 插入排序
func sortInsert(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		tmp := nums[i]

		j := i
		for j > 0 && tmp > nums[j-1] {
			nums[j] = nums[j-1]
			j--
		}

		// 存在比其小的数，插入
		if j != i {
			nums[j] = tmp
		}
	}
	return nums
}

func main() {
	// // 冒泡排序
	// nums := []int{9, 2, 5, 4, 1, 2, 3, 4, 5}
	// nums = sortBubble(nums)
	// fmt.Printf("nums %v \n", nums)

	// nums = []int{1, 2, 3, 4, 5}
	// nums = sortBubble(nums)
	// fmt.Printf("nums %v \n", nums)

	// // 选择排序
	// nums := []int{9, 2, 5, 4, 1, 2, 3, 4, 5}
	// nums = sortSelection(nums)
	// fmt.Printf("nums %v \n", nums)

	// nums = []int{1, 2, 3, 4, 5}
	// nums = sortSelection(nums)
	// fmt.Printf("nums %v \n", nums)

	// 插入排序
	nums := []int{9, 2, 5, 4, 1, 2, 3, 4, 5}
	nums = sortInsert(nums)
	fmt.Printf("nums %v \n", nums)

	nums = []int{1, 2, 3, 4, 5}
	nums = sortInsert(nums)
	fmt.Printf("nums %v \n", nums)
}
