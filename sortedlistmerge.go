package piscine

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	x := n1
	for x.Next != nil {
		x = x.Next
	}
	x.Next = n2
	i := n1
	for ;i != nil; i = i.Next {
		j := i.Next
		for ;j != nil; j = j.Next {
			if i.Data > j.Data {
				i.Data, j.Data = j.Data, i.Data
			}
		}
	}
	return n1
}

