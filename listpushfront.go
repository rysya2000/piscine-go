package piscine

func ListPushFront(l *List, data interface{}) {
	n := &NodeL{Data: data}
	n.Next = l.Head
	l.Head = n
}
