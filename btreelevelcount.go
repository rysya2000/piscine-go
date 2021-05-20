package piscine

func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 1
	}
	return BTreeLevelCount(root.Left) + 1
	return BTreeLevelCount(root.Right) + 1
}
