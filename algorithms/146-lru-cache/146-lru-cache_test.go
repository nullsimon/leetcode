package lrucache

import (
	"testing"
)

func TestLruCache(t *testing.T) {
	lru := Constructor(2)
	lru.Put(1, 1)
	t.Logf(`lru is : %v`, lru)
	lru.Put(2, 2)
	t.Logf(`lru is : %v`, lru)
	lru.Get(1)
	lru.Put(3, 3)
	lru.Get(2)
	lru.Put(4, 4)
	lru.Get(1)
	lru.Get(3)
	lru.Get(4)

	lru2 := Constructor(2)
	lru2.Get(2)
	lru2.Put(2, 6) // {{2,6}}
	lru2.Get(1)
	lru2.Put(1, 5) // {{1,5}, {2,6}}
	lru2.Put(1, 2) // {{1,2}, {2,6}} check if exist, not add
	lru2.Get(1)
	lru2.Get(2)

}
