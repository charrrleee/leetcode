package _00__Same_Tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p != nil && q != nil {
		return isSameTree(p.Left, q.Left) &&
			isSameTree(p.Right, q.Right) &&
			p.Val == q.Val
	}

	return p == q

}
