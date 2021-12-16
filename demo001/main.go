package main

import "fmt"

// 合并两个有序链表

type ListNode struct {
	Val  int       // 链表数据域，任意数据类型
	Next *ListNode // 链表后继指针域，下一节点的内存地址
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

func GenerateListNode(arr []int) *ListNode {
	n1 := &ListNode{Val: 0}
	tmpN := n1
	for _, v := range arr {
		tmpN.Next = &ListNode{Val: v}

		tmpN = tmpN.Next
	}
	return n1.Next
}

func GetListNode1() *ListNode {
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 3}
	n4 := &ListNode{Val: 4}
	n3.Next = n4
	n2.Next = n3
	n1.Next = n2
	return n1
}

func GetListNode2() *ListNode {
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 3}
	n3 := &ListNode{Val: 4}
	n4 := &ListNode{Val: 5}
	n3.Next = n4
	n2.Next = n3
	n1.Next = n2
	return n1
}

func main() {
	// n1 := GetListNode1()
	// nn1 := GetListNode2()

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
	// ln := mergeTwoLists(n1, nn1)
	// for {
	// 	fmt.Printf("%d,", ln.Val)
	// 	if ln.Next == nil {
	// 		break
	// 	}
	// 	ln = ln.Next
	// }

	wn := GenerateListNode([]int{1, 1, 2, 3, 3, 4, 4, 5})
	for {
		fmt.Printf("%d,", wn.Val)
		if wn.Next == nil {
			break
		}
		wn = wn.Next
	}
}
