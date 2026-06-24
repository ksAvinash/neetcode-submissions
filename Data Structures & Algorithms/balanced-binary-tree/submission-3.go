/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func height(root *TreeNode) int {
	if root == nil {
		return -1
	}
	left := height(root.Left)
	right := height(root.Right)
	return max(left, right) + 1
}


func isBalanced(root *TreeNode) bool {
    if root == nil {
		return true
	}
	left := height(root.Left)
	right := height(root.Right)
	diff := math.Abs(float64(left-right))
	return diff <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}
