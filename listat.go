package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	n := 0
	for l != nil {
		n++
		l = l.Next
		if pos == n {
			return l
		}
	}
	return nil
}
