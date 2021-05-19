package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	n := &NodeI{Data: data_ref}
	n.Next = l
	i := n
	j := n.Next
	for j != nil {
		if i.Data > j.Data {
			i.Data, j.Data = j.Data, i.Data
			i = j
		}
		j = j.Next
	}
	return n
}
