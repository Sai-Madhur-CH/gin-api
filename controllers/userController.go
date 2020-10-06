package controllers

import (
	"net/http"
	// "fmt"
	userService "gin-api/services"
	"gin-api/models"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)


// In the controller get the actual request params or request body
// And call the services
// get the service response and handle error if any
// And finally return the service response
func UserRegister(c *gin.Context) {
	user := models.User{} 
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, _ := c.Get("db")
	conn := db.(pgx.Conn)

	userService.UserRegister(c, &conn, user)
	
}