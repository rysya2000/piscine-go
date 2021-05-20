package piscine

func BTreeIsBinary(root *TreeNode) bool {
	if root != nil {
		return true
	}
	return BTreeIsBinary(root.Left) && BTreeIsBinary(root.Right)
	return false
}
