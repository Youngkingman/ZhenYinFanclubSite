package controllers

import (
	algorithm "basic/models/algorithm/searchmethod"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CompareParam struct {
	Dense        float64 `json:"dense" `
	Cols         int     `json:"cols"`
	Rows         int     `json:"rows"`
	Method       string  `json:"method"`
	CostLow      int     `json:"costL"`
	CostHigh     int     `json:"costH"`
	StartPointX  int     `json:"startX"`
	StartPointY  int     `json:"startY"`
	TargetPointX int     `json:"targetX"`
	TargetPointY int     `json:"targetY"`
}

func GetSearchPage(c *gin.Context) {
	c.HTML(http.StatusOK, "view/showsearch.tmpl", gin.H{
		"title": "search_algorithm",
	})
}

func RecordSelectData(c *gin.Context) {
	var formdata CompareParam
	// ShouldBind 和 Bind 类似，不过会在出错时退出而不是返回400状态码
	c.BindJSON(&formdata)
	retData, feasibleMap, costMap := algorithm.Compare(formdata.Rows, formdata.Cols,
		formdata.Dense, formdata.CostLow, formdata.CostHigh,
		[2]int{formdata.StartPointX, formdata.StartPointY},
		[2]int{formdata.TargetPointX, formdata.TargetPointY}, 0,
	)
	retMap := struct {
		RetData     map[string]interface{} `json:"retData"`
		FeasibleMap [][]int                `json:"feasibleMap"`
		CostMap     [][]int                `json:"costMap"`
	}{
		RetData:     retData,
		FeasibleMap: feasibleMap,
		CostMap:     costMap,
	}
	c.JSON(200, retMap)
}
