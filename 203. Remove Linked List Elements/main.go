package _03__Remove_Linked_List_Elements

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {

	pointer := head

	// set nil case
	if head == nil {
		return head
	}

	// check pointer next val
	for pointer.Next != nil {
		if pointer.Next.Val == val {
			pointer.Next = pointer.Next.Next
		} else {
			pointer = pointer.Next
		}
	}

	// check current pointer case
	if head.Val == val {
		return head.Next
	}

	return head
}
