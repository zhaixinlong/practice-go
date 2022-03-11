package list_node_slice

import "github.com/zhaixinlong/practice-go/model"

//  插入排序 改变val
func InsertionSortOnVal(head *model.ListNode) *model.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	lastSorted := head

	for lastSorted != nil {
		current := lastSorted.Next
		for current != nil {
			if lastSorted.Val > current.Val {
				tmpValue := current.Val

				current.Val = lastSorted.Val
				lastSorted.Val = tmpValue
			}
			current = current.Next
		}
		lastSorted = lastSorted.Next
	}
	return head
}

// 插入排序 改变node
func InsertionSort(head *model.ListNode) *model.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummyHead := &model.ListNode{Next: head}
	lastSorted, curr := head, head.Next
	for curr != nil {
		if lastSorted.Val <= curr.Val {
			lastSorted = lastSorted.Next
		} else {
			prev := dummyHead
			for prev.Next.Val <= curr.Val {
				prev = prev.Next
			}

			lastSorted.Next = curr.Next
			curr.Next = prev.Next
			prev.Next = curr
		}
		curr = lastSorted.Next
	}
	return dummyHead.Next
}
