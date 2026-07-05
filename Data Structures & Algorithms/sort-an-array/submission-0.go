type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func newNode(v int) *TreeNode {
	return &TreeNode{
		Val: v,
	}
}

func insert(t *TreeNode, v int) *TreeNode {
	if t == nil {
		return newNode(v)
	}
	if v <= t.Val {
		t.Left = insert(t.Left, v)
	} else {
		t.Right = insert(t.Right, v)
	}
	return t
}

func inOrder(t *TreeNode) []int {
	if t == nil {
		return []int{}
	}

	res := inOrder(t.Left)
	res = append(res, t.Val)
	res = append(res, inOrder(t.Right)...)
	return res
}


func sortArray(nums []int) []int {
	var root *TreeNode
	for _, v := range nums{
		root = insert(root, v)
	}
    return inOrder(root)
}
