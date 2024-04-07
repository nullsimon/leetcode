package timeneededtoinformallemployees

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	return numOfMinutes_version6(n, headID, manager, informTime)
}

// time limit max, faield 28715
func numOfMinutes_version1(n int, headID int, manager []int, informTime []int) int {
	if n == 1 {
		return 0
	}
	var ch = make(chan int, 100000)
	var dist [100000]int
	var visted = make(map[int]bool)
	visted[headID] = true
	dist[headID] = informTime[headID]
	res := 0
	ch <- headID
	for len(ch) > 0 {
		var size = len(ch)
		for i := 0; i < size; i++ {
			cur := <-ch
			for j := 0; j < len(manager); j++ {
				if visted[j] {
					continue
				}
				if manager[j] == cur {
					ch <- j

					visted[j] = true

					dist[j] = dist[cur] + informTime[j]
					if dist[j] > res {
						res = dist[j]
					}
				}
			}
		}
	}
	return res
}

// alter size using map,failed 28715
func numOfMinutes_version2(n int, headID int, manager []int, informTime []int) int {
	if n == 1 {
		return 0
	}
	var ch = make(chan int, 100000)
	var dist [100000]int
	var unvisitedManager = make(map[int]int)
	for i := 0; i < len(manager); i++ {
		unvisitedManager[i] = manager[i]
	}
	delete(unvisitedManager, headID)
	dist[headID] = informTime[headID]
	res := 0
	ch <- headID
	for len(ch) > 0 {
		var size = len(ch)
		for i := 0; i < size; i++ {
			cur := <-ch
			for k := range unvisitedManager {

				if unvisitedManager[k] == cur {
					ch <- k

					delete(unvisitedManager, k)

					dist[k] = dist[cur] + informTime[k]
					if dist[k] > res {
						res = dist[k]
					}
				}
			}
		}
	}
	return res
}

// using slice not channel, guess who fast, failed 59376
func numOfMinutes_version3(n int, headID int, manager []int, informTime []int) int {
	if n == 1 {
		return 0
	}
	var s []int
	var dist [100000]int

	dist[headID] = informTime[headID]
	res := 0
	s = append(s, headID)
	for len(s) > 0 {
		var size = len(s)
		for i := 0; i < size; i++ {
			cur := s[0]
			s = s[1:]
			for j := 0; j < len(manager); j++ {

				if manager[j] == cur {
					s = append(s, j)
					dist[j] = dist[cur] + informTime[j]
					if dist[j] > res {
						res = dist[j]
					}
				}
			}
		}
	}
	return res
}

//change size to map, reduce one loop, but still failed at  91621 n
func numOfMinutes_version4(n int, headID int, manager []int, informTime []int) int {
	if n == 1 {
		return 0
	}
	var m = make(map[int]bool)
	var nm = make(map[int]bool)
	var dist [100000]int
	dist[headID] = informTime[headID]
	res := 0
	m[headID] = true
	for len(m) > 0 {

		for j := 0; j < len(manager); j++ {
			head := manager[j]
			if _, ok := m[head]; ok {
				nm[j] = true

				dist[j] = dist[head] + informTime[j]
				if dist[j] > res {
					res = dist[j]
				}
			}
		}
		m = nm
		nm = make(map[int]bool)
	}
	return res
}

type Node struct {
	Val        int
	informTime int
	Children   []*Node
}

func numOfMinutes_version6(n int, headID int, manager []int, informTime []int) int {
	if n == 1 {
		return 0
	}
	root := new(Node)
	root.Val = headID
	root.informTime = informTime[headID]
	existed := make(map[int]*Node)
	existed[headID] = root

	for i := 0; i < len(manager); i++ {
		if i == headID {
			continue
		}

		node, ok := existed[i]
		if !ok {
			node = new(Node)
			node.Val = i
			node.informTime = informTime[i]
			existed[i] = node
		}
		n, ok := existed[manager[i]]
		if ok {
			//manager node exists, append to it
			n.Children = append(n.Children, node)
		} else {
			nl := new(Node)
			nl.Val = manager[i]
			nl.informTime = informTime[manager[i]]
			nl.Children = append(nl.Children, node)
			existed[manager[i]] = nl
		}
	}

	var dist [100000]int
	dist[headID] = informTime[headID]
	var res = 0
	var ch = make(chan *Node, 100000)
	ch <- root
	for len(ch) > 0 {
		for i := 0; i < len(ch); i++ {
			cur := <-ch
			for j := 0; j < len(cur.Children); j++ {
				c := cur.Children[j]
				ch <- c
				dist[c.Val] = dist[cur.Val] + c.informTime
				if dist[c.Val] > res {
					res = dist[c.Val]
				}
			}
		}
	}
	return res
}
