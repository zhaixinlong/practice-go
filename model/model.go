package model

type ListNode struct {
	Val  int       // 链表数据域，任意数据类型
	Next *ListNode // 链表后继指针域，下一节点的内存地址
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
