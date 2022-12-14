package _3__Remove_Duplicates_from_Sorted_List_

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	tail := head
	for tail.Next != nil {
		if tail.Next.Val == tail.Val {
			tail.Next = tail.Next.Next
		} else {
			tail = tail.Next
		}
	}
	return head
}

func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var out = head
	var hash = make(map[int]bool)
	hash[head.Val] = true

	var cur = out

	for head != nil {
		_, ok := hash[head.Val]

		if !ok {
			cur.Next = head
			hash[head.Val] = true
			cur = cur.Next
		}

		head = head.Next
	}

	cur.Next = nil

	return out
}
