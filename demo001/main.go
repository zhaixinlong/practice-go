package main

import (
	"fmt"

	"github.com/zhaixinlong/practice-go/model"
)

// 合并两个有序链表

// 递归
func mergeTwoLists(list1 *model.ListNode, list2 *model.ListNode) *model.ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	}
	list2.Next = mergeTwoLists(list1, list2.Next)
	return list2
}

// 遍历
func mergeTwoListsTraverse(n1 *model.ListNode, n2 *model.ListNode) *model.ListNode {
	rootNode := &model.ListNode{}

	curr := rootNode
	for n1 != nil || n2 != nil {
		if n1 == nil {
			curr.Next = n2
			curr = curr.Next
			break
		}
		if n2 == nil {
			curr.Next = n1
			curr = curr.Next
			break
		}

		if n1.Val < n2.Val {
			curr.Next = n1
			n1 = n1.Next
		} else {
			curr.Next = n2
			n2 = n2.Next
		}
		curr = curr.Next
	}
	return rootNode.Next
}

func main() {
	n1 := model.GetListNode1()
	nn1 := model.GetListNode2()

	// 遍历实现
	ln := mergeTwoListsTraverse(n1, nn1)
	for {
		fmt.Printf("%d,", ln.Val)
		if ln.Next == nil {
			break
		}
		ln = ln.Next
	}

	// 递归实现
	// ln := mergeTwoLists(n1, nn1)
	// for {
	// 	fmt.Printf("%d,", ln.Val)
	// 	if ln.Next == nil {
	// 		break
	// 	}
	// 	ln = ln.Next
	// }

	// wn := GenerateListNode([]int{1, 1, 2, 3, 3, 4, 4, 5})
	// for {
	// 	fmt.Printf("%d,", wn.Val)
	// 	if wn.Next == nil {
	// 		break
	// 	}
	// 	wn = wn.Next
	// }
}
