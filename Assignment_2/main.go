package main

import (
	"Assignment_2/config"
	"Assignment_2/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.InitDB()
	DBConn := &controllers.DBconn{DB: db}
	router := gin.Default()

	router.POST("/welcome", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome")
	})

	// CRUD - Create, Read, Update & Delete

	//Create
	router.POST("/person", DBConn.CreateOrder)

	//Read -- localhost:8080/person/1 --> Param
	router.GET("/person/:id", DBConn.GetOrderById)
	router.GET("/persons", DBConn.GetOrders)

	//Update -- localhost:8080/person?id=1 --> Query
	router.PUT("/person", DBConn.UpdateOrder)

	//Delete
	router.DELETE("/person/:id", DBConn.DeletePerson)

	router.Run(":1234")
}
