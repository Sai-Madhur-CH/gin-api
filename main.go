package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	routers "gin-api/routers"
)

// Initializing the router globally
var router *gin.Engine


func main() {
	fmt.Println("This is the starting point of the application")
	
	router = gin.Default() // Default router
	
	routers.InitializeRoutes(router) // Initialize all the routers

	router.Run(":8000") // server runs on 8000 port
}