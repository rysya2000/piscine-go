package piscine

func BTreeIsBinary(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left != nil && root.Left.Data >= root.Data {
		return false
	}
	if root.Right != nil && root.Right.Data < root.Data {
		return false
	}
	if BTreeIsBinary(root.Left) == false || BTreeIsBinary(root.Right) == false {
		return false
	}
	return true
}
