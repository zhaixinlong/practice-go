package sort_slice

import "fmt"

func partition(arr []int, low, high int) int {
	index := low - 1
	pivotElement := arr[high]
	for i := low; i < high; i++ {
		fmt.Printf("index %d i %d pivotElement %d low %v high %v \n", index, i, pivotElement, low, high)
		if arr[i] <= pivotElement {
			index += 1
			arr[index], arr[i] = arr[i], arr[index]
		}
	}
	arr[index+1], arr[high] = arr[high], arr[index+1]
	return index + 1
}

// QuickSortRange Sorts the specified range within the array
func QuickSortRange(arr []int, low, high int) {
	if len(arr) <= 1 {
		return
	}

	if low < high {
		pivot := partition(arr, low, high)
		fmt.Printf("pivot %d low %v high %v \n", pivot, low, high)
		QuickSortRange(arr, low, pivot-1)
		QuickSortRange(arr, pivot+1, high)
	}
}

// QuickSort Sorts the entire array
func QuickSort(arr []int) []int {
	QuickSortRange(arr, 0, len(arr)-1)
	return arr
}
