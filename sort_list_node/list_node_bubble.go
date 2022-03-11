package list_node_slice

import "github.com/zhaixinlong/practice-go/model"

// 冒泡排序  改变val
func BubbleSort(head *model.ListNode) *model.ListNode {
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
