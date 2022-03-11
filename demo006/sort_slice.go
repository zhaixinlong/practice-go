package main

import (
	"fmt"

	"github.com/zhaixinlong/practice-go/sort_slice"
)

// 排序 slice

func main() {
	var nums []int

	// // 冒泡排序
	// nums := []int{9, 2, 5, 4, 1, 2, 3, 4, 5}
	// nums = sort_slice.BubbleSort(nums)
	// fmt.Printf("BubbleSort nums %v \n", nums)

	// nums = []int{1, 2, 3, 4, 5}
	// nums = sort_slice.BubbleSort(nums)
	// fmt.Printf("BubbleSort nums %v \n", nums)

	// // 选择排序
	// nums = []int{9, 2, 5, 4, 1, 2, 3, 4, 5}
	// nums = sort_slice.SelectionSort(nums)
	// fmt.Printf("SelectionSort nums %v \n", nums)

	// nums = []int{1, 2, 3, 4, 5}
	// nums = sort_slice.SelectionSort(nums)
	// fmt.Printf("SelectionSort nums %v \n", nums)

	// 插入排序
	nums = []int{9, 2, 5, 4, 1, 2, 3, 4, 5}
	nums = sort_slice.InsertionSort(nums)
	fmt.Printf("InsertionSort nums %v \n", nums)

	nums = []int{1, 2, 3, 4, 5}
	nums = sort_slice.InsertionSort(nums)
	fmt.Printf("InsertionSort nums %v \n", nums)

	// // 快速排序
	// n := []int{1, 8, 10, 5, 4, 3, 2, 1, 100}
	// nums = sort_slice.QuickSort(n)
	// fmt.Printf("QuickSort nums %v \n", nums)
}
