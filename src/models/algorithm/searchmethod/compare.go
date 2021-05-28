package algorithm

import (
	"basic/datatrans"
	"fmt"
	"sync"
	"time"
)

var SaveDataFlag bool = false

/*
	@note:this function uses concurrent calculation to boost the comparision between 3
	different kinds of algorithm. Basically gorotine, semaphore and message quene(channel)
	are used in this example. It serves as a good example for multiple tasks calculation mission.

	@input:`row` and `col` size the map. `dense` is the partial of blocked cells in the map
	`costL` and `costH` range the cost of every step of the agent, `strat` and `end` denotes
	the start point and target point you want to reach.

	@author:
*/
func Compare(row, col int, dense float64, costL, costH int, start, end [2]int, id int) (retData map[string]interface{}, Feasible [][]int, CostMap [][]int) {
	retData = make(map[string]interface{})
	fmt.Println("random map is generating....")
	feasibleMap, retMap := MapGenerator(row, col, dense, costL, costH)
	Feasible, CostMap = feasibleMap, retMap
	filename := time.Now().Format("200601021545")
	if SaveDataFlag {
		datatrans.OutputMat(filename, feasibleMap, retMap, id)
	}

	var wg sync.WaitGroup
	datachan := make(chan map[string]interface{}, 10)

	wg.Add(7)
	go func() {
		tstart := time.Now()
		c1, s1, t1 := DijkstraForGrid(feasibleMap, retMap, start, end)
		timecost := time.Since(tstart)
		fmt.Println("Task of Dijkstra serch is over,the total step is", s1, "and the cost is", c1, " time expired ", timecost)
		data := make(map[string]interface{})
		data["Di"] = 1
		data["cost"] = c1
		data["total"] = s1
		data["tract"] = t1
		if SaveDataFlag {
			datatrans.OutputTract(filename+"_Dijkstra_", t1, id)
		}
		datachan <- data
		wg.Done()
	}()
	go func() {
		tstart := time.Now()
		c2, s2, t2 := AstarSearch(feasibleMap, retMap, start, end, HalmintanDistance)
		timecost := time.Since(tstart)
		fmt.Println("Task of A* search is over, the total step is", s2, "and the cost is", c2, " time expired ", timecost)
		data := make(map[string]interface{})
		data["As"] = 1
		data["cost"] = c2
		data["total"] = s2
		data["tract"] = t2
		if SaveDataFlag {
			datatrans.OutputTract(filename+"_Astar_", t2, id)
		}
		datachan <- data
		wg.Done()
	}()
	go func() {
		tstart := time.Now()
		c2, s2, t2 := AstarSearch(feasibleMap, retMap, start, end, ChebyshevDistance)
		timecost := time.Since(tstart)
		fmt.Println("Task of A*C search is over, the total step is", s2, "and the cost is", c2, " time expired ", timecost)
		data := make(map[string]interface{})
		data["AsC"] = 1
		data["cost"] = c2
		data["total"] = s2
		data["tract"] = t2
		if SaveDataFlag {
			datatrans.OutputTract(filename+"_Astar_", t2, id)
		}
		datachan <- data
		wg.Done()
	}()
	go func() {
		tstart := time.Now()
		c3, s3, t3 := AstarSearchDijkstra(feasibleMap, retMap, start, end, HalmintanDistance)
		timecost := time.Since(tstart)
		fmt.Println("Task of Dijkstra with A* is over, the total step is", s3, "and the cost is", c3, " time expired ", timecost)
		data := make(map[string]interface{})
		data["MOA"] = 1
		data["cost"] = c3
		data["total"] = s3
		data["tract"] = t3
		if SaveDataFlag {
			datatrans.OutputTract(filename+"_DijkstraAstar_", t3, id)
		}
		datachan <- data
		wg.Done()
	}()
	go func() {
		tstart := time.Now()
		c3, s3, t3 := BfsSearch(feasibleMap, retMap, start, end)
		timecost := time.Since(tstart)
		fmt.Println("Task of bfs is over, the total step is", s3, "and the cost is", c3, " time expired ", timecost)
		data := make(map[string]interface{})
		data["BFS"] = 1
		data["cost"] = c3
		data["total"] = s3
		data["tract"] = t3
		if SaveDataFlag {
			datatrans.OutputTract(filename+"_bfs_", t3, id)
		}
		datachan <- data
		wg.Done()
	}()
	go func() {
		tstart := time.Now()
		c3, s3, t3 := JPS(feasibleMap, 10, start, end, _HalmintanDistance)
		timecost := time.Since(tstart)
		fmt.Println("Task of JPS is over, the total step is", s3, "and the cost is", c3, t3, " time expired ", timecost)
		data := make(map[string]interface{})
		data["JPS"] = 1
		data["cost"] = c3
		data["total"] = s3
		data["tract"] = t3
		if SaveDataFlag {
			datatrans.OutputTract(filename+"_DijkstraAstar_", t3, id)
		}
		datachan <- data
		wg.Done()
	}()
	go func() {
		tstart := time.Now()
		c3, s3, t3 := BidirectionAstarDijkstra_Normal(feasibleMap, retMap, start, end, HalmintanDistance)
		timecost := time.Since(tstart)
		fmt.Println("Task of BIMOA* is over, the total step is", s3, "and the cost is", c3, " time expired ", timecost)
		data := make(map[string]interface{})
		data["BIMOA"] = 1
		data["cost"] = c3
		data["total"] = s3
		data["tract"] = t3
		if SaveDataFlag {
			datatrans.OutputTract(filename+"_DijkstraAstar_", t3, id)
		}
		datachan <- data
		wg.Done()
	}()
	fmt.Println("Gorotinues are working")
	wg.Wait()
	fmt.Println("Gorotinues finish tasks")
	close(datachan)
	for v := range datachan {
		//retData = append(retData, v)
		if _, has := v["BFS"]; has {
			retData["BFS"] = v
		}
		if _, has := v["MOA"]; has {
			retData["MOA"] = v
		}
		if _, has := v["As"]; has {
			retData["As"] = v
		}
		if _, has := v["Di"]; has {
			retData["Dijkstra"] = v
		}
		if _, has := v["JPS"]; has {
			retData["JPS"] = v
		}

	}
	return
}
