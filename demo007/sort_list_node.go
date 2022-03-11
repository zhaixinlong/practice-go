package main

import (
	"fmt"

	"github.com/zhaixinlong/practice-go/list_node_slice"
	"github.com/zhaixinlong/practice-go/model"
)

func main() {
	wn := model.GenerateListNode([]int{-1, 9, 7, 3, 1, 2, 8})
	// wn := GenerateListNode([]int{3, 4, 1})
	// for {
	// 	fmt.Printf("%d,", wn.Val)

	// 	if wn.Next == nil {
	// 		break
	// 	}
	// 	wn = wn.Next
	// }
	// fmt.Printf("\n")

	// wn = sortBubble(wn)

	// wn = sortInsertionOnVal(wn)

	wn = list_node_slice.InsertionSort(wn)

	for {
		fmt.Printf("%d,", wn.Val)
		if wn.Next == nil {
			break
		}
		wn = wn.Next
	}
}
