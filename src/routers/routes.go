package routers

import (
	"basic/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// 初始化 Gin 框架默认实例，该实例包含了路由、中间件以及配置信息
	r := gin.Default()
	//加载模板静态资源
	r.Static("/static", "./static")
	//加载模板路由
	r.LoadHTMLGlob("view/*")
	r.GET("/show", controllers.GetSearchPage)
	// Ping 测试路由
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	//初期功能(分组路由就尼玛怪异)
	r.POST("/selectdata", controllers.RecordSelectData)
	v2 := r.Group("/u")
	{
		v2.GET("/test", controllers.GetUserInfo)
	}
	return r
}
