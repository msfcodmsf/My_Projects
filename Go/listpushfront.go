package piscine

func ListPushFront(l *List, data interface{}) {
	nesne := &NodeL{Data: data}

	if l.Head == nil {
		l.Head, l.Tail = nesne, l.Head
	} else {
		nesne.Next, l.Head = l.Head, nesne
	}
}
