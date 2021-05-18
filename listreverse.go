package piscine

func ListReverse(l *List) {
	n := l.Head
	l.Tail = n
	var NewNext *NodeL
	for n != nil {
		n, NewNext, n.Next = n.Next, n, NewNext
//		fmt.Println(NewNext)
	}
	l.Head = NewNext
}
