package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

func ListAt(l *NodeL, pos int) *NodeL {
	n := 0
	for l != nil {
		n++
		l = l.Next
		if pos == n {
			return l.Next
		}
	}
	return nil
}
