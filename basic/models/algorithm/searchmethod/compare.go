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
func Compare(row, col int, dense float64, costL, costH int, start, end [2]int, id int) (retData []interface{}) {
	fmt.Println("random map is generating...")
	feasibleMap, retMap := MapGenerator(row, col, dense, costL, costH)
	filename := time.Now().Format("200601021545")
	if SaveDataFlag {
		datatrans.OutputMat(filename, feasibleMap, retMap, id)
	}

	var wg sync.WaitGroup
	datachan := make(chan map[string]interface{}, 10)

	wg.Add(4)
	go func() {
		c1, s1, t1 := DijkstraForGrid(feasibleMap, retMap, start, end)
		fmt.Println("Task of Dijkstra serch is over,the total step is", s1, "and the cost is", c1)
		data := make(map[string]interface{})
		data["cost of dijkstra"] = c1
		data["total step of dijkstra"] = s1
		data["tract of dijkstra"] = t1
		if SaveDataFlag {
			datatrans.OutputTract(filename+"_Dijkstra_", t1, id)
		}
		datachan <- data
		wg.Done()
	}()
	go func() {
		c2, s2, t2 := AstarSearch(feasibleMap, retMap, start, end, HalmintanDistance)
		fmt.Println("Task of A* search is over, the total step is", s2, "and the cost is", c2)
		data := make(map[string]interface{})
		data["cost of A*"] = c2
		data["total step of A*"] = s2
		data["tract of A*"] = t2
		if SaveDataFlag {
			datatrans.OutputTract(filename+"_Astar_", t2, id)
		}
		datachan <- data
		wg.Done()
	}()
	go func() {
		c3, s3, t3 := AstarSearchDijkstra(feasibleMap, retMap, start, end, HalmintanDistance)
		fmt.Println("Task of Dijkstra with A* is over, the total step is", s3, "and the cost is", c3)
		data := make(map[string]interface{})
		data["cost of DijkstraA*"] = c3
		data["total step of DijkstraA*"] = s3
		data["tract of DijkstraA*"] = t3
		if SaveDataFlag {
			datatrans.OutputTract(filename+"_DijkstraAstar_", t3, id)
		}
		datachan <- data
		wg.Done()
	}()
	go func() {
		c3, s3, t3 := BfsSearch(feasibleMap, retMap, start, end)
		fmt.Println("Task of bfs is over, the total step is", s3, "and the cost is", c3)
		data := make(map[string]interface{})
		data["cost of bFS"] = c3
		data["total step of BFS"] = s3
		data["tract of BFS"] = t3
		if SaveDataFlag {
			datatrans.OutputTract(filename+"_bfs_", t3, id)
		}
		datachan <- data
		wg.Done()
	}()
	fmt.Println("Gorotinues are working")
	wg.Wait()
	fmt.Println("Gorotinues finish tasks")
	close(datachan)
	for v := range datachan {
		retData = append(retData, v)
		//fmt.Println(v)
	}
	return
}
