package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Data > elem {
		return BTreeSearchItem(root.Left, elem)
	}
	if root.Data < elem {
		return BTreeSearchItem(root.Right, elem)
	}
	return root
}
