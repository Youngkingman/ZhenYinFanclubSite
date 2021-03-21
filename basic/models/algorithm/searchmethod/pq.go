package algorithm

//Priority Struct implementation
type Node struct {
	Cost   int64 //current cost for djkstra and fvalue for A* and its variants
	Id     int   //unique identify document for points on a matrix, compressed for grid
	Parent *Node //pointer used to extract the route of search
	Gvalue int64 //used in A*, generally the current cost
	Hvalue int64 //used in A*, generally the estimation of distance between target and current position
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Cost < pq[j].Cost
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Node)
	*pq = append(*pq, item)
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}
