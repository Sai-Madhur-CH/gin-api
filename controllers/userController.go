package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	userService "gin-api/services"
)


func UserRegister(c *gin.Context) {
	result, err := userService.UserRegister()
	if err != nil {
		c.JSON(http.StatusNotFound,gin.H{
			"status":"not found",
		})
	}
	c.JSON(http.StatusOK,result)
}