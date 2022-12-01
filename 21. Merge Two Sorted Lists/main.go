package _1__Merge_Two_Sorted_Lists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var head *ListNode

	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	if list1.Val > list2.Val {
		head = list2
		list2 = list2.Next
	} else {
		head = list1
		list1 = list1.Next
	}

	step := head
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			step.Next = list1
			list1 = list1.Next
		} else {
			step.Next = list2
			list2 = list2.Next
		}
		step = step.Next
	}

	if list1 == nil {
		step.Next = list2
	}

	if list2 == nil {
		step.Next = list1
	}

	return head
}
