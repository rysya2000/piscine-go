package piscine

func ListForEachIf(l *List, f func(*NodeL), cond func(*NodeL) bool) {
	n := l.Head
	for n != nil {
		if cond(n) {
			f(n)
		}
		n = n.Next
	}
}
