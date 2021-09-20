package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	r := gin.Default()
	r.GET("/api/sample", getHelloWorld)
	r.GET("/api/todos", getTodos)

	r.Run(":" + port)
}

func getHelloWorld(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", os.Getenv("CORS_ORIGIN"))

	c.JSON(http.StatusOK, gin.H { "message": "Hello World!" })
}

func getTodos(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", os.Getenv("CORS_ORIGIN"))

	type Todo struct {
		Id int
		Title string
	}

	todos := [...] Todo {
		{ Id: 1, Title: "TODO 1" },
		{ Id: 2, Title: "TODO 2" },
		{ Id: 3, Title: "TODO 3" },
		{ Id: 4, Title: "TODO 4" },
		{ Id: 5, Title: "TODO 5" },
	}

	c.JSON(http.StatusOK, gin.H { "todos": todos })
}
