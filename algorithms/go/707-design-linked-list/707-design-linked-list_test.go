package designlinkedlist

import (
	"testing"
)

func TestLinkedList(t *testing.T) {
	l := ConstructorSingly()
	l.AddAtHead(1)
	l.AddAtTail(3)
	l.AddAtIndex(1, 2)
	l.Get(1)
	l.DeleteAtIndex(1)
	l.Get(1)

}
