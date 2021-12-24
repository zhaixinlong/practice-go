package main

import "fmt"

func returnNewSlice(slice []int) ([]int, int) {
	var idx = 1 // 游标，判断比较到了哪个下标

	// 第一位（不存在重复） 默认跳过，不需要重新写入
	for i := 0; i+1 < len(slice); i++ { // i+1 不能 大于数组长度，防止越界
		if slice[i] != slice[i+1] {
			slice[idx] = slice[i+1]
			idx++ //附近两个值不一致，下标++
		}
	}
	return slice[:idx], len(slice)
}

// 有序数组，原地删除重复项 返回非重复数量，以及新数组
func main() {
	slice := []int{0, 0, 1, 2, 3, 4, 4, 5, 6, 7, 7, 8, 9, 10, 10, 10, 11}
	newSlice, sliceSize := returnNewSlice(slice)
	fmt.Printf("new slice is %v original slice len %d \n", newSlice, sliceSize)
}
