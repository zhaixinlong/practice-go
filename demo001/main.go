package main

import "fmt"

// 合并两个有序链表

type ListNode struct {
	Val  int
	Next *ListNode
}

// 递归
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
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
func mergeTwoListsTraverse(list1 *ListNode, list2 *ListNode) *ListNode {
	tmpNode := &ListNode{}

	prev := tmpNode
	for {
		if list1 == nil && list2 == nil {
			break
		}

		if list1 == nil {
			prev.Next = list2
			prev = prev.Next
			break
		}
		if list2 == nil {
			prev.Next = list1
			prev = prev.Next
			break
		}

		if list1.Val < list2.Val {
			prev.Next = list1
			list1 = list1.Next
		} else {
			prev.Next = list2
			list2 = list2.Next
		}

		prev = prev.Next
	}
	return tmpNode.Next
}

func main() {
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 3}
	n4 := &ListNode{Val: 4}
	n3.Next = n4
	n2.Next = n3
	n1.Next = n2

	nn1 := &ListNode{Val: 1}
	nn2 := &ListNode{Val: 3}
	nn3 := &ListNode{Val: 4}
	nn4 := &ListNode{Val: 5}
	nn3.Next = nn4
	nn2.Next = nn3
	nn1.Next = nn2

	// 遍历实现
	// ln := mergeTwoListsTraverse(n1, nn1)
	// for {
	// 	fmt.Printf("%d,", ln.Val)
	// 	if ln.Next == nil {
	// 		break
	// 	}
	// 	ln = ln.Next
	// }

	// 递归实现
	ln := mergeTwoLists(n1, nn1)
	for {
		fmt.Printf("%d,", ln.Val)
		if ln.Next == nil {
			break
		}
		ln = ln.Next
	}
}