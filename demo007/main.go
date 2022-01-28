package main

import (
	"fmt"

	"github.com/zhaixinlong/practice-go/model"
)

func GenerateListNode(arr []int) *model.ListNode {
	n1 := &model.ListNode{Val: 0}
	tmpN := n1
	for _, v := range arr {
		tmpN.Next = &model.ListNode{Val: v}

		tmpN = tmpN.Next
	}
	return n1.Next
}

// 冒泡排序  改变val
func sortBubble(head *model.ListNode) *model.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	isChange := true
	var p *model.ListNode
	for p != head.Next && isChange {
		q := head

		isChange = false
		for q.Next != nil && q.Next != p {
			if q.Val > q.Next.Val {
				val := q.Next.Val

				q.Next.Val = q.Val
				q.Val = val

				isChange = true
			}
			q = q.Next
		}
		p = q
	}
	return head
}

//  插入排序 改变val
func sortInsertionOnVal(head *model.ListNode) *model.ListNode {
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
func sortInsertion(head *model.ListNode) *model.ListNode {
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

func main() {
	wn := GenerateListNode([]int{-1, 9, 7, 3, 1, 2, 8})
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

	wn = sortInsertion(wn)

	for {
		fmt.Printf("%d,", wn.Val)
		if wn.Next == nil {
			break
		}
		wn = wn.Next
	}
}
