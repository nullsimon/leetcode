package designlinkedlist

type Node struct {
	Key   int
	Value int
	Prev  *Node
	Next  *Node
}

type MyLinkedList struct {
	Size int
	Head *Node
	Tail *Node
}

func Constructor() MyLinkedList {
	l := MyLinkedList{
		Size: 0,
		Head: nil,
		Tail: nil,
	}
	return l

}

func (this *MyLinkedList) Get(index int) int {
	return 0
}

func (this *MyLinkedList) AddAtHead(val int) {

}

func (this *MyLinkedList) AddAtTail(val int) {

}

func (this *MyLinkedList) AddAtIndex(index int, val int) {

}

func (this *MyLinkedList) DeleteAtIndex(index int) {

}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

type SinglyNode struct {
	Value int
	Next  *SinglyNode
}

type SinglyLinkedList struct {
	Size int
	Head *SinglyNode
	Tail *SinglyNode
}

func ConstructorSingly() SinglyLinkedList {
	l := SinglyLinkedList{
		Size: 0,
		Head: nil,
		Tail: nil,
	}
	return l
}

func (this *SinglyLinkedList) Get(index int) int {
	if index >= this.Size {
		return -1
	}
	//todo binary search
	n := 0
	l := this.Head
	for l != nil {
		if n == index {
			return l.Value
		}
		l = l.Next
		n++
	}
	return -1
}

func (this *SinglyLinkedList) AddAtHead(val int) {
	sn := &SinglyNode{
		Value: val,
	}

	if this.Size == 0 {
		this.Head = sn
		this.Tail = sn
		this.Size++
		return
	}

	sn.Next = this.Head
	this.Head = sn
	this.Size++
}

func (this *SinglyLinkedList) AddAtTail(val int) {
	sn := &SinglyNode{
		Value: val,
	}

	if this.Size == 0 {
		this.Head = sn
		this.Tail = sn
		this.Size++
		return
	}

	this.Tail.Next = sn
	this.Tail = sn
	this.Size++
}

func (this *SinglyLinkedList) AddAtIndex(index int, val int) {
	if index > this.Size {
		return
	}
	if index == this.Size {
		this.AddAtTail(val)
		return
	}
	if index == 0 {
		this.AddAtHead(val)
		return
	}
	n := 0
	l := this.Head
	for l != nil {
		if n == index-1 {
			addNode := &SinglyNode{
				Value: val,
			}
			next := l.Next
			l.Next = addNode
			addNode.Next = next
			this.Size++
			return
		}
		l = l.Next
		n++
	}
}

func (this *SinglyLinkedList) DeleteAtIndex(index int) {
	if index >= this.Size {
		return
	}
	if this.Size == 1 {
		this.Head = nil
		this.Tail = nil
		this.Size--
		return
	}
	if index == 0 {
		this.Head = this.Head.Next
		this.Size--
		return
	}

	n := 0
	l := this.Head
	preNode := l
	for l != nil {
		if n == index {
			//find it

			preNode.Next = l.Next
			if l.Next == nil {
				this.Tail = preNode
			}
			this.Size--
		}
		preNode = l
		l = l.Next
		n++
	}

}
