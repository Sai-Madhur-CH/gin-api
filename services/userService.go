package services

import (
	"fmt"
	"net/http"
	"gin-api/models"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

// Main bussiness logic is implemented in the service
func UserRegister(c *gin.Context, conn *pgx.Conn,user models.User ) {

	err := user.Register(conn)
	if err != nil {
		fmt.Println("Error in user.Register()")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := user.GetAuthToken()
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"token": token,
			"user_id": user.User_id,
			"name": user.Name,
			"email": user.Email,
		})
		return
	}

}