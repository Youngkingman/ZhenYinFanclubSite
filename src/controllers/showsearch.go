package controllers

import (
	algorithm "basic/models/algorithm/searchmethod"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CompareParam struct {
	Dense        float64 `form:"dense"`
	Cols         int     `form:"cols"`
	Rows         int     `form:"rows"`
	Method       string  `form:"method"`
	CostLow      int     `form:"costL"`
	CostHigh     int     `form:"costH"`
	StartPointX  int     `form:"startX"`
	StartPointY  int     `form:"startY"`
	TargetPointX int     `form:"targetX"`
	TargetPointY int     `form:"targetX"`
}

func GetSearchPage(c *gin.Context) {
	c.HTML(http.StatusOK, "view/showsearch.tmpl", gin.H{
		"title": "search_algorithm",
	})
}

func RecordSelectData(c *gin.Context) {
	var formdata CompareParam
	// ShouldBind 和 Bind 类似，不过会在出错时退出而不是返回400状态码
	c.ShouldBind(&formdata)
	retData := algorithm.Compare(formdata.Rows, formdata.Cols,
		formdata.Dense, formdata.CostLow, formdata.CostHigh,
		[2]int{formdata.StartPointX, formdata.StartPointY},
		[2]int{formdata.TargetPointX, formdata.TargetPointY}, 0,
	)
	c.JSON(200, retData)
}
