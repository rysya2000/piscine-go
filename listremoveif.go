package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	for l.Head != nil && l.Head.Data == data_ref {
		l.Head = l.Head.Next
	}
	n := l.Head
	for n != nil && n.Next != nil {
		if n.Next.Data == data_ref {
			n.Next = n.Next.Next
		} else {
			n = n.Next
		}
	}
}
