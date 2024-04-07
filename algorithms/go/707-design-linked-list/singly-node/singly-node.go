package singlynode

type Node struct {
	Val  int
	Next *Node
}

type LinkedList struct {
	Size int
	Head *Node
}

func Constructor() LinkedList {
	l := LinkedList{
		Size: 0,
		Head: &Node{},
	}
	return l
}

func (this *LinkedList) Get(index int) int {
	if index < 0 || index >= this.Size {
		return -1
	}
	prev := this.Head //dummy node
	for i := 0; i < index+1; i++ {
		prev = prev.Next
	}
	return prev.Val
}

func (this *LinkedList) AddAtHead(val int) {
	this.AddAtIndex(0, val)
}

func (this *LinkedList) AddAtTail(val int) {
	this.AddAtIndex(this.Size, val)
}

func (this *LinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > this.Size {
		return
	}
	prev := this.Head
	for i := 0; i < index; i++ {
		prev = prev.Next
	}
	tmp := &Node{
		Val: val,
	}
	tmp.Next = prev.Next
	prev.Next = tmp
	this.Size++
}

func (this *LinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.Size {
		return
	}
	prev := this.Head
	for i := 0; i < index; i++ {
		prev = prev.Next
	}
	prev.Next = prev.Next.Next
	this.Size--
}
