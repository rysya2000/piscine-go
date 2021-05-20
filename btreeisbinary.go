package piscine

func BTreeIsBinary(root *TreeNode) bool {
	if root.Left.Data > root.Right.Data {
		return false
	}else {
		return true
	}
	return BTreeIsBinary(root.Left) || BTreeIsBinary(root.Right)
}
