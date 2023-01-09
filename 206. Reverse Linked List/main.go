package _06__Reverse_Linked_List

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var node *ListNode

	// 1 > 2 > 3
	for head != nil {

		// 1, nil > 2, (1, node)
		list := &ListNode{Val: head.Val, Next: node}

		// node = 1, node, 2, (1, node)
		node = list

		// push
		head = head.Next

	}
	return node
}
