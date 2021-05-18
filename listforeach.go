package piscine

func ListForEach(l *List, f func(*NodeL)) {
	n := l.Head
	for n != nil {
		f(n)
		n = n.Next
	}
}
