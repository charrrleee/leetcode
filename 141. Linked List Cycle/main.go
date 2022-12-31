package _41__Linked_List_Cycle

type ListNode struct {
	Val  int
	Next *ListNode
}

// this function is check the linklist has cycle or not
// so we use fast and slow to find out fast node and slow node is
func hasCycle(head *ListNode) bool {
	var slow = head
	var fast = head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		/* meeting point in the cycle */
		if slow == fast {
			return true
		}
	}

	return false
}
