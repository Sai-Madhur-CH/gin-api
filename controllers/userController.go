package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	userService "gin-api/services"
)


// In the controller get the actual request params or request body
// And call the services
// get the service response and handle error if any
// And finally return the service response
func UserRegister(c *gin.Context) {
	result, err := userService.UserRegister()
	if err != nil {
		c.JSON(http.StatusNotFound,gin.H{
			"status":"not found",
		})
	}
	c.JSON(http.StatusOK,result)
}