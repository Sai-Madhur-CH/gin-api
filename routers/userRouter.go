package routers

import (
	"github.com/gin-gonic/gin"
	userController "gin-api/controllers"
)


// Get the router of type *gin.Engine form the main.go file and initilize all the routers 
// And point them to the specific controllers
func InitializeRoutes(router *gin.Engine) {
	router.POST("/user",userController.UserRegister	)
}