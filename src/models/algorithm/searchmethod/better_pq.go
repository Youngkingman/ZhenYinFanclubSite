package algorithm

type Motion struct {
	Dealt_X    int
	Dealt_Y    int
	Dealt_Cost float32
}

//Priority Struct implementation
type _Node struct {
	Cost      int
	X         int
	Y         int
	Gvalue    int
	Hvalue    int
	Parent    *_Node
	ForceList []Motion
}

type _PriorityQueue []*_Node

func (pq _PriorityQueue) Len() int { return len(pq) }
func (pq _PriorityQueue) Less(i, j int) bool {
	return pq[i].Cost < pq[j].Cost
}
func (pq _PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *_PriorityQueue) Push(x interface{}) {
	item := x.(*_Node)
	*pq = append(*pq, item)
}
func (pq *_PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func (node _Node) IsEqual(other _Node) bool {
	if node.X == other.X && node.Y == other.Y {
		return true
	}
	return false
}
