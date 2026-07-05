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

func inOrder(t *TreeNode, res *[]int) {
	if t == nil {
		return
	}
	inOrder(t.Left, res)
	*res = append(*res, t.Val)
	inOrder(t.Right, res)
}

func sortArray(nums []int) []int {
	var root *TreeNode
	for _, v := range nums{
		root = insert(root, v)
	}
	
	res := make([]int, 0, len(nums))
	inOrder(root, &res)
    return res
}
