package lrucache

type LRUCache struct {
	Data map[int]int
	S    []int
	Cap  int
}

func Constructor(capacity int) LRUCache {
	lru := new(LRUCache)
	lru.Data = make(map[int]int)
	lru.S = make([]int, capacity)
	lru.Cap = capacity
	return *lru
}

func (this *LRUCache) Get(key int) int {
	if len(this.Data) == 0 {
		return -1
	}
	//put it front
	value, ok := this.Data[key]
	if !ok {
		return -1
	}

	for i := 0; i < len(this.S); i++ {
		if this.S[i] == key {
			var tmp = append(this.S[0:i], this.S[i+1:]...)
			var tmp1 = []int{key}
			this.S = append(tmp1, tmp...)
		}
	}
	return value
}

func (this *LRUCache) Put(key int, value int) {

	exists := this.Get(key)
	this.Data[key] = value
	if exists != -1 {
		return
	}

	if len(this.S) == this.Cap {
		oldkey := this.S[this.Cap-1]
		delete(this.Data, oldkey)
		this.S = this.S[:this.Cap-1]

	}
	var tmp = []int{key}
	this.S = append(tmp, this.S...)

}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
