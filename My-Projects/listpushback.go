package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(list *List, data interface{}) {
	node := &NodeL{Data: data}

	if list.Head == nil {
		list.Head = node
	} else {
		list.Tail.Next = node
	}
	list.Tail = node
}
