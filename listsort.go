package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	if l == nil {
		return l
	}
	n := l
	for ; n != nil; n = n.Next {
		m := n.Next
		for ; m != nil; m = m.Next {
			if n.Data > m.Data {
				n.Data, m.Data = m.Data, n.Data
			}
		}
	}
	return l
}
