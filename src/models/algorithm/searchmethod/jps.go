package algorithm

import (
	"container/heap"
)

const (
	// 未被占据的结点
	Unoccupied = iota
	// 被占据的结点
	Occupied
	// 在容器中的结点
	Exploring
	// 已扩展的结点(已经从容器中弹出的结点)
	Explored
	// 路徑上的结点
	Road
)

//JPS A*(8 directions), good for balance map
//更新節點的時候沒有把東西放進去，比較麻煩，但是光出數據沒問題
func JPS(feasibleMap [][]int, costVert int, start [2]int, target [2]int, hvalue func(int, int, int, int) int) (fcost int, step int, tract [][2]int) {
	//地图副本
	feasibleG := make([][]int, 0)
	for _, v := range feasibleMap {
		arr := make([]int, len(feasibleMap[0]))
		for k, _ := range v {
			arr[k] = v[k]
		}
		feasibleG = append(feasibleG, arr)
	}
	openList := make(_PriorityQueue, 0)
	m := len(feasibleG)
	if m == 0 {
		return 0, 0, nil
	}
	n := len(feasibleG[0])

	feasibleG[target[0]][target[1]] = 0 //somewhat gurantee the feasible
	//部分内置函數
	getForceNeigborList := func(current _Node, v Motion, jumpFlag int) (forceList []Motion) {
		if jumpFlag == 0 { //带判断的垂直跳跃
			if current.Y+1 < n && current.X+v.Dealt_X >= 0 && current.X+v.Dealt_X < m &&
				feasibleG[current.X][current.Y+1] == int(Occupied) && feasibleG[current.X+v.Dealt_X][current.Y+1] == int(Unoccupied) {
				forceList = append(forceList, Motion{v.Dealt_X, 1, 1.414})
			}
			if current.Y-1 >= 0 && current.X+v.Dealt_X >= 0 && current.X+v.Dealt_X < m &&
				feasibleG[current.X][current.Y-1] == int(Occupied) && feasibleG[current.X+v.Dealt_X][current.Y-1] == int(Unoccupied) {
				forceList = append(forceList, Motion{v.Dealt_X, -1, 1.414})
			}
		}
		if jumpFlag == 1 { //带判断的水平跳跃
			if current.X+1 < m && current.Y+v.Dealt_Y >= 0 && current.Y+v.Dealt_Y < n &&
				feasibleG[current.X+1][current.Y] == int(Occupied) && feasibleG[current.X+1][current.Y+v.Dealt_Y] == int(Unoccupied) {
				forceList = append(forceList, Motion{1, v.Dealt_Y, 1.414})
			}
			if current.X-1 >= 0 && current.Y+v.Dealt_Y >= 0 && current.Y+v.Dealt_Y < n &&
				feasibleG[current.X-1][current.Y] == int(Occupied) && feasibleG[current.X-1][current.Y+v.Dealt_Y] == int(Unoccupied) {
				forceList = append(forceList, Motion{-1, v.Dealt_Y, 1.414})
			}
		}
		return
	}
	getNodeStatus := func(this _Node) int {
		//在此處越界判斷真是妙啊
		if this.X < 0 || this.Y < 0 || this.X >= m || this.Y >= n {
			return int(Occupied)
		} else {
			return feasibleG[this.X][this.Y]
		}
	}

	motionList := []Motion{
		{-1, 1, 1.414},
		{1, 1, 1.414},
		{1, -1, 1.414},
		{-1, -1, 1.414},
	}
	initItem := _Node{
		Cost:      0,
		X:         start[0],
		Y:         start[1],
		Parent:    nil,
		Gvalue:    0,
		Hvalue:    hvalue(start[0], start[1], target[0], target[1]),
		ForceList: motionList,
	}
	heap.Push(&openList, &initItem)

	flag := false
	//A*框架開始
	for len(openList) > 0 {
		if flag {
			break
		}

		relaxPoint := heap.Pop(&openList).(*_Node)
		step++

		for _, v := range relaxPoint.ForceList {
			if flag {
				break
			}
			p := _Node{
				Cost:   relaxPoint.Gvalue + hvalue(relaxPoint.X, relaxPoint.Y, target[0], target[1]),
				X:      relaxPoint.X,
				Y:      relaxPoint.Y,
				Parent: relaxPoint.Parent,
				Gvalue: relaxPoint.Gvalue,
				Hvalue: relaxPoint.Hvalue,
			}
			for getNodeStatus(p) != int(Occupied) {
				if flag {
					break
				}
				//垂直跳跃
				current := _Node{
					Cost:   p.Cost,
					X:      p.X,
					Y:      p.Y,
					Parent: p.Parent,
					Gvalue: p.Gvalue,
					Hvalue: p.Hvalue,
				}
				for getNodeStatus(current) != int(Occupied) {
					if getNodeStatus(current) == int(Exploring) {
						current.X = current.X + v.Dealt_X
						current.Gvalue = current.Gvalue + costVert
						current.Parent = &p
						current.Cost = current.Gvalue + hvalue(current.X, current.Y, target[0], target[1])
						continue
					}
					forceList := getForceNeigborList(current, v, 0)
					if len(forceList) != 0 {
						current.ForceList = forceList
						if !p.IsEqual(*relaxPoint) && !p.IsEqual(current) {
							feasibleG[p.X][p.Y] = int(Exploring)
						}
						feasibleG[current.X][current.Y] = int(Exploring)
						heap.Push(&openList, &current)
						break
					}
					if getNodeStatus(current) == int(Unoccupied) {
						feasibleG[current.X][current.Y] = int(Explored)
					}
					current.X = current.X + v.Dealt_X
					current.Gvalue = current.Gvalue + costVert
					current.Parent = &p
					current.Cost = current.Gvalue + hvalue(current.X, current.Y, target[0], target[1])

					//獲取目標，開始提取路徑
					if current.X == target[0] && current.Y == target[1] {
						feasibleG[p.X][p.Y] = int(Exploring)
						fcost = current.Cost
						for current.Parent != nil {
							tract = append(tract, [2]int{current.X, current.Y})
							current = *current.Parent
						}
						return
					}
				}
				//水平跳躍
				current = _Node{
					Cost:   p.Cost,
					X:      p.X,
					Y:      p.Y,
					Parent: p.Parent,
					Gvalue: p.Gvalue,
					Hvalue: p.Hvalue,
				}
				for getNodeStatus(current) != int(Occupied) {
					if getNodeStatus(current) == int(Exploring) {
						current.Y = current.Y + v.Dealt_Y
						current.Gvalue = current.Gvalue + costVert
						current.Parent = &p
						current.Cost = current.Gvalue + hvalue(current.X, current.Y, target[0], target[1])
						continue
					}
					forceList := getForceNeigborList(current, v, 0)
					if len(forceList) != 0 {
						current.ForceList = forceList
						if !p.IsEqual(*relaxPoint) && !p.IsEqual(current) {
							feasibleG[p.X][p.Y] = int(Exploring)
						}
						feasibleG[current.X][current.Y] = int(Exploring)
						heap.Push(&openList, &current)
						break
					}
					if getNodeStatus(current) == int(Unoccupied) {
						feasibleG[current.X][current.Y] = int(Explored)
					}
					current.Y = current.Y + v.Dealt_Y
					current.Gvalue = current.Gvalue + costVert
					current.Parent = &p
					current.Cost = current.Gvalue + hvalue(current.X, current.Y, target[0], target[1])

					//獲取目標，開始提取路徑
					if current.X == target[0] && current.Y == target[1] {
						feasibleG[p.X][p.Y] = int(Exploring)
						fcost = current.Cost
						for current.Parent != nil {
							tract = append(tract, [2]int{current.X, current.Y})
							current = *current.Parent
						}
						return
					}
				}
				//對角綫跳躍更新
				p.X, p.Y, p.Gvalue = p.X+v.Dealt_X, p.Y+v.Dealt_Y, int(1.414*float64(costVert))
				p.Parent = relaxPoint
				p.Cost = p.Gvalue + hvalue(p.X, p.Y, target[0], target[1])
			}
		}
	}
	return
}

//Alternative in function JPS, the F value evaluation method
func _HalmintanDistance(currentX, currentY, targetX, targetY int) int {
	abs1 := currentX - targetX
	if abs1 < 0 {
		abs1 = -abs1
	}
	abs2 := currentY - targetY
	if abs2 < 0 {
		abs2 = -abs2
	}
	return abs1 + abs2
}

//Alternative in function JPS, the F value evaluation method
func _ChebyshevDistance(currentX, currentY, targetX, targetY int) int {
	abs1 := currentX - targetX
	if abs1 < 0 {
		abs1 = -abs1
	}
	abs2 := currentY - targetY
	if abs2 < 0 {
		abs2 = -abs2
	}
	return int(float64(abs1+abs2) + 1.414213*float64(min(int64(abs1), int64(abs2))))
}
