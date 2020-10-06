package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"context"
	routers "gin-api/routers"
)

// Initializing the router globally
var router *gin.Engine


func main() {
	fmt.Println("This is the starting point of the application")

	conn, err := connectDB() // For connect to postgresql
	if err != nil{
		return
	}

	router = gin.Default() // Default router

	router.Use(dbMiddleware(*conn)) // To add this db connection to every router
	
	routers.InitializeRoutes(router) // Initialize all the routers

	router.Run(":8000") // server runs on 8000 port
}


// Using pgx to connec to the postgresql db
// This function will return connection on success of type pgx.Conn
// If there is a error while connecting the db it throws the error
// Reference https://godoc.org/github.com/jackc/pgx
func connectDB() (c *pgx.Conn, err error) {
	conn, err := pgx.Connect(context.Background(),
				 "postgres://postgres:postgres@127.0.0.1:5431/gin-db")
	if err != nil {
		fmt.Println("Error while connecting db")
		fmt.Println(err.Error())
	}

	_ = conn.Ping(context.Background())

	return conn, err
}


// This function is used to add the created connection in the Context 
// so that its avaliable all over the application
func dbMiddleware(conn pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context)  {
		c.Set("db", conn)
		c.Next()
	}
}