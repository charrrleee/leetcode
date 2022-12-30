package _4__Binary_Tree_Inorder_Traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//Morris
func inorderTraversal(root *TreeNode) []int {
	var list []int
	if root == nil {
		return list
	}

	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	parentNode := root
	for parentNode != nil {
		if parentNode.Left == nil {
			list = append(list, parentNode.Val)
			parentNode = parentNode.Right
			continue
		}

		currentNode := parentNode.Left

		for currentNode.Right != nil {
			currentNode = currentNode.Right
		}

		currentNode.Right = parentNode

		parentNode, parentNode.Left = parentNode.Left, nil
	}

	return list
}

func inorderTraversal1(root *TreeNode) []int {

	list, currentNode := []int{}, root

	if currentNode != nil {
		list = append(list, inorderTraversal1(currentNode.Left)...)
		list = append(list, currentNode.Val)
		list = append(list, inorderTraversal1(currentNode.Right)...)
	}

	return list

}

func inorderTraversal3(root *TreeNode) []int {

	list, currentNode := []int{}, root
	var stack []*TreeNode

	for currentNode != nil || len(stack) != 0 {
		for currentNode != nil {
			stack = append(stack, currentNode)
			currentNode = currentNode.Left
		}

		currentNode = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		list = append(list, currentNode.Val)
		currentNode = currentNode.Right
	}
	return list
}
