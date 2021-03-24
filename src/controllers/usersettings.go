package controllers

import (
	"basic/models/users"

	"github.com/gin-gonic/gin"
)

func GetUserInfo(c *gin.Context) {
	userinfo := users.GetUserInfo(1)
	c.JSON(200, userinfo)
}
