package main

import (
	"github.com/gin-gonic/gin"
	"lola.com/routes"
	"lola.com/globals"
	_ "github.com/lib/pq"

)


const (
	dbname	  = "restaurant"
	user	  = "postgres"
	password  = "Aa123456"
	host 	  = "localhost"
	port      = 5432
)	


func main() {
	router := gin.Default()
	globals.InitializeConnectionPool()
	defer globals.ConnectionPool.Close()
	
	router.GET("/",func(c *gin.Context) {
		c.JSON(200,gin.H{})
	})
	routes.Orders(router.Group("/orders"))
	router.Run("0.0.0.0:8080")

	
}
