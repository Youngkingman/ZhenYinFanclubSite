package algorithm

import (
	"container/heap"
	"math"
	"math/rand"
	"sync"
	"time"
)

//MAX_INT
const max = int64(^uint64(0) >> 1)

//ExampleMatrix
func NewMatrix() (ret [][]int) {
	m := 200000000
	l0 := make([]int, 8)
	l0[0], l0[1], l0[2], l0[3], l0[4], l0[5], l0[6], l0[7] = 0, 7, 3, m, m, m, m, m
	ret = append(ret, l0)
	l1 := make([]int, 8)
	l1[0], l1[1], l1[2], l1[3], l1[4], l1[5], l1[6], l1[7] = 7, 0, m, 4, m, m, m, 9
	ret = append(ret, l1)
	l2 := make([]int, 8)
	l2[0], l2[1], l2[2], l2[3], l2[4], l2[5], l2[6], l2[7] = 3, m, 0, 5, m, m, m, m
	ret = append(ret, l2)
	l3 := make([]int, 8)
	l3[0], l3[1], l3[2], l3[3], l3[4], l3[5], l3[6], l3[7] = m, 4, 5, 0, 5, m, 2, m
	ret = append(ret, l3)
	l4 := make([]int, 8)
	l4[0], l4[1], l4[2], l4[3], l4[4], l4[5], l4[6], l4[7] = m, m, m, 5, 0, 11, m, m
	ret = append(ret, l4)
	l5 := make([]int, 8)
	l5[0], l5[1], l5[2], l5[3], l5[4], l5[5], l5[6], l5[7] = m, m, m, m, 11, 0, 10, m
	ret = append(ret, l5)
	l6 := make([]int, 8)
	l6[0], l6[1], l6[2], l6[3], l6[4], l6[5], l6[6], l6[7] = m, m, m, 2, m, 10, 0, 2
	ret = append(ret, l6)
	l7 := make([]int, 8)
	l7[0], l7[1], l7[2], l7[3], l7[4], l7[5], l7[6], l7[7] = m, 9, m, m, m, m, 2, 0
	ret = append(ret, l7)
	return
}

/*
	@note:This function serves as a mapgenerator which generating test environment
	for search algorithm.

	@input:Parameter `m` and `n` denote the row number and col number of the grid.
	Parameter `dense` denotes the blocked cell of grid. `costL` and `costH` range
	the open low boundary and close upper boundary of the process of cost generation.

	@output:`retFeasibel` returns the map of a task, and `retCost` returns the cost for an
	angent moving to the current cell from its cell neigbors.
*/
func MapGenerator(m int, n int, dense float64, costL int, costH int) (retFeasible [][]int, retCost [][]int) {
	totalCount := m * n
	//Feasible Matrix
	blockCount := int(float64(m*n) * dense)
	bucketA := make([]int, totalCount)
	for i := 0; i < blockCount; i++ {
		bucketA[i] = 1
	}
	for i := blockCount; i < totalCount; i++ {
		bucketA[i] = 0
	}

	knuthDurstenfeldShuffle(bucketA, 10)
	for i := 0; i < m; i++ {
		arr := make([]int, n)
		for j, _ := range arr {
			arr[j] = bucketA[0]
			bucketA = bucketA[1:]
		}
		retFeasible = append(retFeasible, arr)
	}
	//Cost Matrix
	bucketA = make([]int, totalCount)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i, _ := range bucketA {
		bucketA[i] = costL + r.Intn(costH-costL)
	}
	for i := 0; i < m; i++ {
		arr := make([]int, n)
		for j, _ := range arr {
			arr[j] = bucketA[0]
			bucketA = bucketA[1:]
		}
		retCost = append(retCost, arr)
	}
	return
}

/*
	@note:This function using Dijkstra Algorithm to find the shortest path in
	a graph

	@input: `graph` is the neibor matrix of every vertex of graph. `start`
	is the vertex index of start vertex and `target` is the vertex index of target vertex
s
	@output:the shortest cost to reach the target
*/
func DijkstraForNeiborMat(graph [][]int, start int, target int) int64 {
	if len(graph) < 1 {
		return 0
	}
	pointCount := len(graph)
	setA := make(PriorityQueue, 0)
	dp := make([]int64, pointCount)
	relax := make([]bool, pointCount)
	for k, _ := range dp {
		dp[k] = max
	}
	dp[start] = 0
	initItem := &Node{
		Cost:   0,
		Id:     start,
		Parent: nil,
	}
	heap.Push(&setA, initItem)
	for len(setA) > 0 {
		relaxPoint := heap.Pop(&setA).(*Node)
		id := relaxPoint.Id
		relax[id] = true
		for k, v := range graph[id] {
			if int64(v) < max && v > 0 && relax[k] == false {
				cost := min(int64(v)+dp[id], dp[k])
				dp[k] = cost
				item := Node{
					Cost:   cost,
					Id:     k,
					Parent: relaxPoint,
				}
				heap.Push(&setA, &item)
			}
		}
	}
	return dp[target]
}

/*
	@note:Here DON'T USE THIS FUNCTION.IT'S EXTREMELY SLOW AND IT KILL THE MEMORY!!!
	depth-first search to find out the route. Not always have the least cost,

	@input:

	@output:
*/
func DfsSearch(feasibleMap [][]int, costG [][]int, start [2]int, target [2]int) (retCost int64, step int64, tract [][2]int) {
	feasibleG := make([][]int, 0)
	for _, v := range feasibleMap {
		arr := make([]int, len(feasibleMap[0]))
		for k, _ := range v {
			arr[k] = v[k]
		}
		feasibleG = append(feasibleG, arr)
	}
	m := len(feasibleG)
	if m == 0 {
		return
	}
	n := len(feasibleG[0])
	initItem := &Node{
		Cost:   0,
		Id:     start[0]*n + start[1],
		Parent: nil,
	}
	rowNeigbor, colNeigbor := []int{1, -1}, []int{1, -1}
	//closure
	var search func(item *Node)
	search = func(item *Node) {
		nowRow, nowCol := item.Id/n, item.Id%n
		step++
		if nowRow == target[0] && nowCol == target[1] {
			tractCompress := extract(item)
			tract = decodeTract(tractCompress, n)
			retCost = item.Cost
			return
		}
		for _, v := range rowNeigbor {
			if nowRow+v >= 0 && nowRow+v < m && feasibleG[nowRow+v][nowCol] == 0 {
				feasibleG[nowRow+v][nowCol] = 2
				newItem := &Node{
					Cost:   item.Cost + int64(costG[nowRow+v][nowCol]),
					Id:     (nowRow+v)*n + nowCol,
					Parent: item,
				}
				search(newItem)
			}
		}
		for _, v := range colNeigbor {
			if nowCol+v >= 0 && nowCol+v < n && feasibleG[nowRow+v][nowCol] == 0 {
				feasibleG[nowRow][nowCol+v] = 2
				newItem := &Node{
					Cost:   item.Cost + int64(costG[nowRow][nowCol+v]),
					Id:     nowRow*n + nowCol + v,
					Parent: item,
				}
				search(newItem)
			}
		}
	}
	search(initItem)
	return
}

/*
	@note:Breadth-first search.Basically the fastest method, but the cost is extremely high and not always
	find out a best route. But it performs better than dijkstra in equal-cost grid for it doesn't need any
	priority quene operation.

	@input:
	@output:
*/
func BfsSearch(feasibleMap [][]int, costG [][]int, start [2]int, target [2]int) (retCost int64, step int64, tract [][2]int) {
	feasibleG := make([][]int, 0)
	for _, v := range feasibleMap {
		arr := make([]int, len(feasibleMap[0]))
		for k, _ := range v {
			arr[k] = v[k]
		}
		feasibleG = append(feasibleG, arr)
	}
	m := len(feasibleG)
	if m == 0 {
		return
	}
	n := len(feasibleG[0])
	initItem := Node{
		Cost:   0,
		Id:     start[0]*n + start[1],
		Parent: nil,
	}
	normalQuene := make([]*Node, 0)
	normalQuene = append(normalQuene, &initItem)
	rowNeigbor, colNeigbor := []int{1, -1}, []int{1, -1}
	for len(normalQuene) > 0 {
		relaxPoint := normalQuene[0]
		normalQuene = normalQuene[1:]
		step++
		id := relaxPoint.Id
		nowRow, nowCol := id/n, id%n
		if nowRow == target[0] && nowCol == target[1] {
			retCost = relaxPoint.Cost
			tractCompress := extract(relaxPoint)
			tract = decodeTract(tractCompress, n)
			return
		}
		for _, v1 := range rowNeigbor {
			tempX, tempY := nowRow+v1, nowCol
			//inside the boundary & unblocked & unrelaxed
			if tempX < m && tempX >= 0 && tempY < n && tempY >= 0 && feasibleG[tempX][tempY] == 0 {
				cost := relaxPoint.Cost + int64(costG[tempX][tempY])
				feasibleG[tempX][tempY] = 2
				item := Node{
					Cost:   cost,
					Id:     tempX*n + tempY,
					Parent: relaxPoint,
				}
				normalQuene = append(normalQuene, &item)
			}
		}
		for _, v1 := range colNeigbor {
			tempX, tempY := nowRow, nowCol+v1
			//inside the boundary & unblocked & unrelaxed
			if tempX < m && tempX >= 0 && tempY < n && tempY >= 0 && feasibleG[tempX][tempY] == 0 {
				cost := relaxPoint.Cost + int64(costG[tempX][tempY])
				feasibleG[tempX][tempY] = 2
				item := Node{
					Cost:   cost,
					Id:     tempX*n + tempY,
					Parent: relaxPoint,
				}
				normalQuene = append(normalQuene, &item)
			}
		}
	}
	return
}

/*
	@note:Here feasibleG is a two dimension slice with elements
	zeros and ones denote unblock cells and block cells respectively,
	and costG in the same type denote the current G cost in A* Algorithem.
	Furthermore, when element of feasibleG become 2, it means that the point
	have been relaxed by djkstra algorithem. It returns the least cost and
	least cost for points relaxed(dynamic programming matrix).

	@input:

	@output:
*/
func DijkstraForGrid(feasibleMap [][]int, costG [][]int, start [2]int, target [2]int) (retCost int64, step int64, tract [][2]int) {
	feasibleG := make([][]int, 0)
	for _, v := range feasibleMap {
		arr := make([]int, len(feasibleMap[0]))
		for k, _ := range v {
			arr[k] = v[k]
		}
		feasibleG = append(feasibleG, arr)
	}
	setA := make(PriorityQueue, 0)
	m := len(feasibleG)
	if m == 0 {
		return
	}
	n := len(feasibleG[0])
	dpG := make([][]int64, 0)
	//init the dp value
	for i := 0; i < m; i++ {
		arr := make([]int64, n)
		for j, _ := range arr {
			arr[j] = int64(max)
		}
		dpG = append(dpG, arr)
	}
	dpG[start[0]][start[1]] = 0
	feasibleG[target[0]][target[1]] = 0 //somewhat gurantee the feasible
	initItem := Node{
		Cost:   0,
		Id:     start[0]*n + start[1],
		Parent: nil,
	}
	heap.Push(&setA, &initItem)
	//BFS for djkstra, 8 connected
	//rowNeigbor, colNeigbor := []int{0, 1, -1}, []int{0, 1, -1}
	rowNeigbor, colNeigbor := []int{1, -1}, []int{1, -1}
	for len(setA) > 0 {
		relaxPoint := heap.Pop(&setA).(*Node)
		step++
		id := relaxPoint.Id
		nowRow, nowCol := id/n, id%n
		if feasibleG[nowRow][nowCol] == 2 {
			continue
		}
		feasibleG[nowRow][nowCol] = 2

		if nowRow == target[0] && nowCol == target[1] {
			retCost = dpG[nowRow][nowCol]
			tractCompress := extract(relaxPoint)
			tract = decodeTract(tractCompress, n)
			return
		}
		for _, v1 := range rowNeigbor {
			tempX, tempY := nowRow+v1, nowCol
			//inside the boundary & unblocked & unrelaxed
			if tempX < m && tempX >= 0 && tempY < n && tempY >= 0 && feasibleG[tempX][tempY] == 0 {
				cost := min(dpG[nowRow][nowCol]+int64(costG[tempX][tempY]), dpG[tempX][tempY])
				dpG[tempX][tempY] = cost
				item := Node{
					Cost:   cost,
					Id:     tempX*n + tempY,
					Parent: relaxPoint,
				}

				heap.Push(&setA, &item)
			}
		}
		for _, v1 := range colNeigbor {
			tempX, tempY := nowRow, nowCol+v1
			//inside the boundary & unblocked & unrelaxed
			if tempX < m && tempX >= 0 && tempY < n && tempY >= 0 && feasibleG[tempX][tempY] == 0 {
				cost := min(dpG[nowRow][nowCol]+int64(costG[tempX][tempY]), dpG[tempX][tempY])
				dpG[tempX][tempY] = cost
				item := Node{
					Cost:   cost,
					Id:     tempX*n + tempY,
					Parent: relaxPoint,
				}
				heap.Push(&setA, &item)
			}
		}
	}
	return
}

/*
	@note:A* has extra interface as a function hvalue to calculate the h-value.
	As for hvalue, it requires four parameters denote current row index, current col index,
	target row index and target col index respectively which return the h value. Other
	parameters share the same fucntion as DjkstraForGrid.

	@input:

	@output:
*/
func AstarSearch(feasibleMap [][]int, costG [][]int, start [2]int, target [2]int, hvalue func(int, int, int, int) int64) (fcost int64, step int64, tract [][2]int) {
	feasibleG := make([][]int, 0)
	for _, v := range feasibleMap {
		arr := make([]int, len(feasibleMap[0]))
		for k, _ := range v {
			arr[k] = v[k]
		}
		feasibleG = append(feasibleG, arr)
	}
	setA := make(PriorityQueue, 0)
	m := len(feasibleG)
	if m == 0 {
		return 0, 0, nil
	}
	n := len(feasibleG[0])
	dpG := make([][]int64, 0)
	//init the dp value
	for i := 0; i < m; i++ {
		arr := make([]int64, n)
		for j, _ := range arr {
			arr[j] = max
		}
		dpG = append(dpG, arr)
	}
	dpG[start[0]][start[1]] = 0
	feasibleG[target[0]][target[1]] = 0 //somewhat gurantee the feasible
	initItem := Node{
		Cost:   0,
		Id:     start[0]*n + start[1],
		Parent: nil,
		Gvalue: 0,
		Hvalue: hvalue(start[0], start[1], target[0], target[1]),
	}
	heap.Push(&setA, &initItem)
	//BFS for djkstra
	//rowNeigbor, colNeigbor := []int{0, 1, -1}, []int{0, 1, -1} //8-connected,remains to be construct
	rowNeigbor, colNeigbor := []int{1, -1}, []int{1, -1} //4-connected
	for len(setA) > 0 {
		relaxPoint := heap.Pop(&setA).(*Node)
		step++
		id := relaxPoint.Id
		nowRow, nowCol := id/n, id%n

		if nowRow == target[0] && nowCol == target[1] {
			fcost = dpG[nowRow][nowCol]
			tractCompress := extract(relaxPoint)
			tract = decodeTract(tractCompress, n)
			return
		}
		for _, v1 := range rowNeigbor {
			tempX, tempY := nowRow+v1, nowCol
			//inside the boundary & unblocked & unrelaxed
			if tempX < m && tempX >= 0 && tempY < n && tempY >= 0 && feasibleG[tempX][tempY] == 0 {
				feasibleG[tempX][tempY] = 2
				//gval := min(dpG[nowRow][nowCol]+costG[tempX][tempY], dpG[tempX][tempY])
				gval := dpG[nowRow][nowCol] + int64(costG[tempX][tempY])
				hval := hvalue(tempX, tempY, target[0], target[1])
				dpG[tempX][tempY] = gval
				item := Node{
					Cost:   gval + hval,
					Id:     tempX*n + tempY,
					Parent: relaxPoint,
					Gvalue: gval,
					Hvalue: hval,
				}
				heap.Push(&setA, &item)
			}
		}
		for _, v1 := range colNeigbor {
			tempX, tempY := nowRow, nowCol+v1
			//inside the boundary & unblocked & unrelaxed
			if tempX < m && tempX >= 0 && tempY < n && tempY >= 0 && feasibleG[tempX][tempY] == 0 {
				feasibleG[tempX][tempY] = 2
				//gval := min(dpG[nowRow][nowCol]+costG[tempX][tempY], dpG[tempX][tempY])
				gval := dpG[nowRow][nowCol] + int64(costG[tempX][tempY])
				hval := hvalue(tempX, tempY, target[0], target[1])
				dpG[tempX][tempY] = gval
				item := Node{
					Cost:   gval + hval,
					Id:     tempX*n + tempY,
					Parent: relaxPoint,
					Gvalue: gval,
					Hvalue: hval,
				}
				heap.Push(&setA, &item)
			}
		}

	}
	return
}

/*
	@note:Improve the A* search by using cost in Dijgkstra Algorithm to make sure the
	path shortest and using less step than Dijkstra Algoritm. It also have the same interface
	as AstarSearch.

	@input:

	@output:
*/
func AstarSearchDijkstra(feasibleMap [][]int, costG [][]int, start [2]int, target [2]int, hvalue func(int, int, int, int) int64) (fcost int64, step int64, tract [][2]int) {
	feasibleG := make([][]int, 0)
	for _, v := range feasibleMap {
		arr := make([]int, len(feasibleMap[0]))
		for k, _ := range v {
			arr[k] = v[k]
		}
		feasibleG = append(feasibleG, arr)
	}
	setA := make(PriorityQueue, 0)
	m := len(feasibleG)
	if m == 0 {
		return 0, 0, nil
	}
	n := len(feasibleG[0])
	dpG := make([][]int64, 0)
	//init the dp value
	for i := 0; i < m; i++ {
		arr := make([]int64, n)
		for j, _ := range arr {
			arr[j] = max
		}
		dpG = append(dpG, arr)
	}
	dpG[start[0]][start[1]] = 0
	feasibleG[target[0]][target[1]] = 0 //somewhat gurantee the feasible
	initItem := Node{
		Cost:   0,
		Id:     start[0]*n + start[1],
		Parent: nil,
		Gvalue: 0,
		Hvalue: hvalue(start[0], start[1], target[0], target[1]),
	}
	heap.Push(&setA, &initItem)
	//BFS for djkstra
	//rowNeigbor, colNeigbor := []int{0, 1, -1}, []int{0, 1, -1} //8-connected,remains to be construct
	rowNeigbor, colNeigbor := []int{1, -1}, []int{1, -1} //4-connected
	for len(setA) > 0 {
		relaxPoint := heap.Pop(&setA).(*Node)
		step++
		id := relaxPoint.Id
		nowRow, nowCol := id/n, id%n
		if feasibleG[nowRow][nowCol] == 2 {
			continue
		}
		feasibleG[nowRow][nowCol] = 2

		if nowRow == target[0] && nowCol == target[1] {
			fcost = dpG[nowRow][nowCol]
			tractCompress := extract(relaxPoint)
			tract = decodeTract(tractCompress, n)
			return
		}
		for _, v1 := range rowNeigbor {
			tempX, tempY := nowRow+v1, nowCol
			//inside the boundary & unblocked & unrelaxed
			if tempX < m && tempX >= 0 && tempY < n && tempY >= 0 && feasibleG[tempX][tempY] == 0 {
				gval := min(dpG[nowRow][nowCol]+int64(costG[tempX][tempY]), dpG[tempX][tempY])
				//gval := dpG[nowRow][nowCol] + costG[tempX][tempY]
				hval := hvalue(tempX, tempY, target[0], target[1])
				dpG[tempX][tempY] = gval
				item := Node{
					Cost:   gval + hval,
					Id:     tempX*n + tempY,
					Parent: relaxPoint,
					Gvalue: gval,
					Hvalue: hval,
				}
				heap.Push(&setA, &item)
			}
		}
		for _, v1 := range colNeigbor {
			tempX, tempY := nowRow, nowCol+v1
			//inside the boundary & unblocked & unrelaxed
			if tempX < m && tempX >= 0 && tempY < n && tempY >= 0 && feasibleG[tempX][tempY] == 0 {
				gval := min(dpG[nowRow][nowCol]+int64(costG[tempX][tempY]), dpG[tempX][tempY])
				//gval := dpG[nowRow][nowCol] + costG[tempX][tempY]
				hval := hvalue(tempX, tempY, target[0], target[1])
				dpG[tempX][tempY] = gval
				item := Node{
					Cost:   gval + hval,
					Id:     tempX*n + tempY,
					Parent: relaxPoint,
					Gvalue: gval,
					Hvalue: hval,
				}
				heap.Push(&setA, &item)
			}
		}
	}
	return
}

/*
	@note: This function use concurency method to search the route in forward and backward direction,
	it's useful in task such as agents to find each other.

	@input:

	@output:
*/
func BidirectionAstarDijkstra(feasibleMap [][]int, costG [][]int, start [2]int, target [2]int, hvalue func(int, int, int, int) int64) (fcost int64, step int64, tract [][2]int) {
	//initialization
	feasibleG := make([][]int, 0) //should be threadsafe
	for _, v := range feasibleMap {
		arr := make([]int, len(feasibleMap[0]))
		for k, _ := range v {
			arr[k] = v[k]
		}
		feasibleG = append(feasibleG, arr)
	}
	setA := make(PriorityQueue, 0)
	setB := make(PriorityQueue, 0)
	m := len(feasibleG)
	if m == 0 {
		return 0, 0, nil
	}
	n := len(feasibleG[0])
	dpG := make([][]int64, 0) //should be threadsafe
	dpG[start[0]][start[1]] = 0
	//init the dp value
	for i := 0; i < m; i++ {
		arr := make([]int64, n)
		for j, _ := range arr {
			arr[j] = max
		}
		dpG = append(dpG, arr)
	}

	var wg sync.WaitGroup //semaphore
	var mux *sync.Mutex   //mutex

	initItemStart := Node{
		Cost:   0,
		Id:     start[0]*n + start[1],
		Parent: nil,
		Gvalue: 0,
		Hvalue: hvalue(start[0], start[1], target[0], target[1]),
	}
	heap.Push(&setA, &initItemStart)
	initItemTarget := Node{
		Cost:   int64(costG[target[0]][target[1]]),
		Id:     target[0]*n + target[1],
		Parent: nil,
		Gvalue: 0,
		Hvalue: hvalue(start[0], start[1], target[0], target[1]),
	}
	heap.Push(&setB, &initItemTarget)

	//multi-thread work
	wg.Add(2)
	stepA := 0
	stopChan := make(chan int) //communication channel for  threads
	dataChan := make(chan *Node)

	//forward search (unfinished)
	go func() {
		rowNeigbor, colNeigbor := []int{1, -1}, []int{1, -1} //4-connected

	loop:
		for len(setA) > 0 {
			//lock here to make sure
			mux.Lock()

			//getting message from backword search
			select {
			case <-stopChan:
				//series of operation, remains to be construct
				break loop
			default:
			}

			relaxPoint := heap.Pop(&setA).(*Node)
			stepA++
			id := relaxPoint.Id
			nowRow, nowCol := id/n, id%n
			/*--------Critical Section-------*/
			if feasibleG[nowRow][nowCol] == 2 {
				mux.Unlock()
				continue
			}
			if feasibleG[nowRow][nowCol] == 3 {
				stopChan <- 0
				dataChan <- relaxPoint
				close(dataChan)
				mux.Unlock()
				/*
					when the forward search meet the backward search
					stop every thing
				*/

				return
			}
			feasibleG[nowRow][nowCol] = 2
			if nowRow == target[0] && nowCol == target[1] {
				fcost = dpG[nowRow][nowCol]
				tractCompress := extract(relaxPoint)
				tract = decodeTract(tractCompress, n)
				mux.Unlock()
				return
			}

			for _, v1 := range rowNeigbor {
				tempX, tempY := nowRow+v1, nowCol
				//inside the boundary & unblocked & unrelaxed
				if tempX < m && tempX >= 0 && tempY < n && tempY >= 0 && feasibleG[tempX][tempY] == 0 {
					gval := min(dpG[nowRow][nowCol]+int64(costG[tempX][tempY]), dpG[tempX][tempY])
					//gval := dpG[nowRow][nowCol] + costG[tempX][tempY]
					hval := hvalue(tempX, tempY, target[0], target[1])
					dpG[tempX][tempY] = gval
					item := Node{
						Cost:   gval + hval,
						Id:     tempX*n + tempY,
						Parent: relaxPoint,
						Gvalue: gval,
						Hvalue: hval,
					}

					heap.Push(&setA, &item)
				}
			}
			for _, v1 := range colNeigbor {
				tempX, tempY := nowRow, nowCol+v1
				//inside the boundary & unblocked & unrelaxed
				if tempX < m && tempX >= 0 && tempY < n && tempY >= 0 && feasibleG[tempX][tempY] == 0 {
					gval := min(dpG[nowRow][nowCol]+int64(costG[tempX][tempY]), dpG[tempX][tempY])
					//gval := dpG[nowRow][nowCol] + costG[tempX][tempY]
					hval := hvalue(tempX, tempY, target[0], target[1])
					dpG[tempX][tempY] = gval
					item := Node{
						Cost:   gval + hval,
						Id:     tempX*n + tempY,
						Parent: relaxPoint,
						Gvalue: gval,
						Hvalue: hval,
					}
					heap.Push(&setA, &item)
				}
			}
			mux.Unlock()
			/*-------critical section end-----------*/
		}
		wg.Done()
	}()

	//backward search (unfinished)
	go func() {
		wg.Done()
	}()

	wg.Wait()
	return
}

//no concurreny version
func BidirectionAstarDijkstra_Normal(feasibleMap [][]int, costG [][]int, start [2]int, target [2]int, hvalue func(int, int, int, int) int64) (fcost int64, step int64, tract [][2]int) {
	feasibleG := make([][]int, 0)
	for _, v := range feasibleMap {
		arr := make([]int, len(feasibleMap[0]))
		for k, _ := range v {
			arr[k] = v[k]
		}
		feasibleG = append(feasibleG, arr)
	}
	setA := make(PriorityQueue, 0)
	setB := make(PriorityQueue, 0)
	relaxedA := make([]Node, 0)
	relaxedB := make([]Node, 0)
	m := len(feasibleG)
	if m == 0 {
		return 0, 0, nil
	}
	n := len(feasibleG[0])
	dpG := make([][]int64, 0)
	//init the dp value
	for i := 0; i < m; i++ {
		arr := make([]int64, n)
		for j, _ := range arr {
			arr[j] = max
		}
		dpG = append(dpG, arr)
	}
	dpG[start[0]][start[1]] = 0
	dpG[target[0]][target[1]] = 0
	feasibleG[target[0]][target[1]] = 0 //somewhat gurantee the feasible
	feasibleG[start[0]][start[1]] = 0
	initItem := Node{
		Cost:   0,
		Id:     start[0]*n + start[1],
		Parent: nil,
		Gvalue: 0,
		Hvalue: hvalue(start[0], start[1], target[0], target[1]),
	}
	initTarget := Node{
		Cost:   0,
		Id:     target[0]*n + target[1],
		Parent: nil,
		Gvalue: 0,
		Hvalue: hvalue(start[0], start[1], target[0], target[1]),
	}
	//for dpG[target[0]][target[1]]=0, it should be add to the final
	//initCost := int64(costG[target[0]][target[1]])
	heap.Push(&setA, &initItem)
	heap.Push(&setB, &initTarget)
	stepA, stepB := 0, 0
	//BFS for djkstra
	//rowNeigbor, colNeigbor := []int{0, 1, -1}, []int{0, 1, -1} //8-connected,remains to be construct
	rowNeigbor, colNeigbor := []int{1, -1}, []int{1, -1} //4-connected
	for len(setA) > 0 && len(setB) > 0 {
		relaxPointA := heap.Pop(&setA).(*Node)
		relaxedA = append(relaxedA, *relaxPointA)
		stepA++
		idA := relaxPointA.Id
		nowRowA, nowColA := idA/n, idA%n
		if feasibleG[nowRowA][nowColA] == 2 && len(setB) > 0 {
			goto DEALB
		}
		if len(setB) <= 0 {
			continue
		}
		if feasibleG[nowRowA][nowColA] == 3 {
			for len(setB) > 0 {
				temp := heap.Pop(&setB).(*Node)
				if temp.Id == idA {
					fcost = dpG[nowRowA][nowColA] + int64(costG[target[0]][target[1]]) + int64(relaxPointA.Parent.Gvalue)
					step = int64(stepA) + int64(stepB)
					tractA := decodeTract(extract(relaxPointA), n)
					tractB := decodeTract(extract(temp), n)
					reverse(tractB)
					tract = append(tract, tractB...)
					tract = append(tract, tractA[1:]...)
					return
				}
			}
			for _, temp := range relaxedB {
				if temp.Id == idA {
					fcost = dpG[nowRowA][nowColA] + int64(costG[target[0]][target[1]]) + int64(relaxPointA.Parent.Gvalue)
					step = int64(stepA) + int64(stepB)
					tractA := decodeTract(extract(relaxPointA), n)
					tractB := decodeTract(extract(&temp), n)
					reverse(tractB)
					tract = append(tract, tractB...)
					tract = append(tract, tractA[1:]...)
					return
				}
			}
			return
		}
		feasibleG[nowRowA][nowColA] = 2

		if nowRowA == target[0] && nowColA == target[1] {
			fcost = dpG[nowRowA][nowColA]
			step = int64(stepA) + int64(stepB)
			tractCompress := extract(relaxPointA)
			tract = decodeTract(tractCompress, n)
			return
		}
		for _, v1 := range rowNeigbor {
			tempX, tempY := nowRowA+v1, nowColA
			//inside the boundary & unblocked & unrelaxed
			if tempX < m && tempX >= 0 && tempY < n && tempY >= 0 && feasibleG[tempX][tempY] == 0 {
				gval := min(dpG[nowRowA][nowColA]+int64(costG[tempX][tempY]), dpG[tempX][tempY])
				//gval := dpG[nowRow][nowCol] + costG[tempX][tempY]
				hval := hvalue(tempX, tempY, target[0], target[1])
				dpG[tempX][tempY] = gval
				item := Node{
					Cost:   gval + hval,
					Id:     tempX*n + tempY,
					Parent: relaxPointA,
					Gvalue: gval,
					Hvalue: hval,
				}

				heap.Push(&setA, &item)
			}
		}
		for _, v1 := range colNeigbor {
			tempX, tempY := nowRowA, nowColA+v1
			//inside the boundary & unblocked & unrelaxed
			if tempX < m && tempX >= 0 && tempY < n && tempY >= 0 && feasibleG[tempX][tempY] == 0 {
				gval := min(dpG[nowRowA][nowColA]+int64(costG[tempX][tempY]), dpG[tempX][tempY])
				//gval := dpG[nowRow][nowCol] + costG[tempX][tempY]
				hval := hvalue(tempX, tempY, target[0], target[1])
				dpG[tempX][tempY] = gval
				item := Node{
					Cost:   gval + hval,
					Id:     tempX*n + tempY,
					Parent: relaxPointA,
					Gvalue: gval,
					Hvalue: hval,
				}
				heap.Push(&setA, &item)
			}
		}
	DEALB:
		relaxPointB := heap.Pop(&setB).(*Node)
		relaxedB = append(relaxedB, *relaxPointB)
		stepB++
		idB := relaxPointB.Id
		nowRowB, nowColB := idB/n, idB%n
		if feasibleG[nowRowB][nowColB] == 3 {
			continue
		}
		if feasibleG[nowRowB][nowColB] == 2 {
			for len(setA) > 0 {
				temp := heap.Pop(&setA).(*Node)
				if temp.Id == idB {
					fcost = dpG[nowRowB][nowColB] + int64(costG[target[0]][target[1]]) + int64(relaxPointB.Parent.Gvalue)
					step = int64(stepA) + int64(stepB)
					tractB := decodeTract(extract(relaxPointB), n)
					tractA := decodeTract(extract(temp), n)
					reverse(tractB)
					tract = append(tract, tractB...)
					tract = append(tract, tractA[1:]...)
					return
				}
			}
			for _, temp := range relaxedA {
				if temp.Id == idB {
					fcost = dpG[nowRowB][nowColB] + int64(costG[target[0]][target[1]]) + int64(relaxPointB.Parent.Gvalue)
					step = int64(stepA) + int64(stepB)
					tractB := decodeTract(extract(relaxPointB), n)
					tractA := decodeTract(extract(&temp), n)
					reverse(tractB)
					tract = append(tract, tractB...)
					tract = append(tract, tractA[1:]...)
					return
				}
			}
			return
		}
		feasibleG[nowRowB][nowColB] = 3

		if nowRowB == start[0] && nowColB == start[1] {
			fcost = dpG[nowRowB][nowColB] + int64(costG[target[0]][target[1]]-costG[start[0]][start[1]])
			step = int64(stepA) + int64(stepB)
			tractCompress := extract(relaxPointB)
			tract = decodeTract(tractCompress, n)
			return
		}
		for _, v1 := range rowNeigbor {
			tempX, tempY := nowRowB+v1, nowColB
			//inside the boundary & unblocked & unrelaxed
			if tempX < m && tempX >= 0 && tempY < n && tempY >= 0 && feasibleG[tempX][tempY] == 0 {
				gval := min(dpG[nowRowB][nowColB]+int64(costG[tempX][tempY]), dpG[tempX][tempY])
				//gval := dpG[nowRow][nowCol] + costG[tempX][tempY]
				hval := hvalue(tempX, tempY, start[0], start[1])
				dpG[tempX][tempY] = gval
				item := Node{
					Cost:   gval + hval,
					Id:     tempX*n + tempY,
					Parent: relaxPointB,
					Gvalue: gval,
					Hvalue: hval,
				}

				heap.Push(&setB, &item)
			}
		}
		for _, v1 := range colNeigbor {
			tempX, tempY := nowRowB, nowColB+v1
			//inside the boundary & unblocked & unrelaxed
			if tempX < m && tempX >= 0 && tempY < n && tempY >= 0 && feasibleG[tempX][tempY] == 0 {
				gval := min(dpG[nowRowB][nowColB]+int64(costG[tempX][tempY]), dpG[tempX][tempY])
				//gval := dpG[nowRow][nowCol] + costG[tempX][tempY]
				hval := hvalue(tempX, tempY, start[0], start[1])
				dpG[tempX][tempY] = gval
				item := Node{
					Cost:   gval + hval,
					Id:     tempX*n + tempY,
					Parent: relaxPointB,
					Gvalue: gval,
					Hvalue: hval,
				}
				heap.Push(&setB, &item)
			}
		}
	}
	return
}

//Alternative in function A*, the F value evaluation method
func HalmintanDistance(currentX, currentY, targetX, targetY int) int64 {
	abs1 := currentX - targetX
	if abs1 < 0 {
		abs1 = -abs1
	}
	abs2 := currentY - targetY
	if abs2 < 0 {
		abs2 = -abs2
	}
	return int64(abs1 + abs2)
}

//Alternative in function A*, the F value evaluation method
func EulerDistance(currentX, currentY, targetX, targetY int) int64 {
	abs1 := currentX - targetY
	abs2 := targetX - currentY
	return int64(math.Sqrt(float64(abs1*abs1 + abs2*abs2)))
}

//Alternative in function A*, the F value evaluation method
func ChebyshevDistance(currentX, currentY, targetX, targetY int) int64 {
	abs1 := currentX - targetX
	if abs1 < 0 {
		abs1 = -abs1
	}
	abs2 := currentY - targetY
	if abs2 < 0 {
		abs2 = -abs2
	}
	return int64(float64(abs1+abs2) + 1.414213*float64(min(int64(abs1), int64(abs2))))
}

//Shuffling data for random tasks
func knuthDurstenfeldShuffle(list []int, additionNum int) {
	l := len(list)
	if l <= 1 {
		return
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano() + int64(additionNum<<1)))
	for l > 0 {
		randIdx := r.Intn(l)
		list[l-1], list[randIdx] = list[randIdx], list[l-1]
		l--
	}
}

//util
func min(i, j int64) int64 {
	if i < j {
		return i
	} else {
		return j
	}
}

//imply to extract the route after calculate
func extract(target *Node) []int {
	ret := make([]int, 0)
	for target != nil {
		ret = append(ret, target.Id)
		target = target.Parent
	}
	return ret
}

//reverse extract
func reverseExtract(target *Node) []int {
	temp := Node{}
	var reverse func(target *Node)
	reverse = func(target *Node) {
		if target.Parent == nil {
			temp = *target
			target.Parent.Parent = target
			target.Parent = nil
			return
		}
		reverse(target.Parent)
		target.Parent.Parent = target
		target.Parent = nil
	}
	reverse(target)
	ret := make([]int, 0)
	for &temp != nil {
		ret = append(ret, temp.Id)
		temp = *temp.Parent
	}
	return ret
}

//imply to decode the compress route of algorithm
func decodeTract(tract []int, n int) (ret [][2]int) {
	for _, v := range tract {
		temp := [2]int{v / n, v % n}
		ret = append(ret, temp)
	}
	return
}

//reverse the tract(seems there is a better way to reverse the listnode)
func reverse(tract [][2]int) {
	l := len(tract)
	for i := 0; i < l/2; i++ {
		tract[i], tract[l-1-i] = tract[l-1-i], tract[i]
	}
}

//calculate cost from tract
func calculatCost(tract [][2]int, costG [][]int) (ret int64) {
	for _, val := range tract {
		ret += int64(costG[val[0]][val[1]])
	}
	return
}
