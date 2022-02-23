package main

import (
	"net/http"
	"os"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	port := os.Getenv("PORT")

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{os.Getenv("CORS_ORIGIN")},
		AllowMethods: []string{"PUT", "PATCH", "DELETE"},
	}))

	r.GET("/api/sample", getHelloWorld)
	r.GET("/api/todos", getTodos)
	r.POST("/api/todos", addTodo)
	r.PATCH("/api/todos/:ID", updateTodo)
	r.DELETE("/api/todos/:ID", deleteTodo)

	r.Run(":" + port)
}

type Todo struct {
	gorm.Model
	Title string
}

func gormConnect() *gorm.DB {
	dsn := os.Getenv("DATABASE_URL")
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&Todo{})

	return db
}

func getHelloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H { "message": "Hello World!" })
}

func getTodos(c *gin.Context) {
	var todos [] Todo

	db := gormConnect()

	db.Find(&todos)

	c.JSON(http.StatusOK, gin.H { "todos": todos })
}

func addTodo(c *gin.Context) {
	type InputTodo struct {
		Title string
	}

	var inputTodo InputTodo
	c.BindJSON(&inputTodo)

	newTodo := Todo { Title: inputTodo.Title }

	db := gormConnect()

	db.Create(&newTodo)

	c.JSON(http.StatusOK, gin.H { "todo": newTodo })
}

func updateTodo(c *gin.Context) {
	type InputTodo struct {
		Title string
	}

	Id := c.Param("ID")
	id, _ := strconv.Atoi(Id)

	db := gormConnect()

	var todo Todo

	db.First(&todo, id)

	var inputTodo InputTodo
	c.BindJSON(&inputTodo)

	db.Model(&todo).Update("Title", inputTodo.Title)

	c.JSON(http.StatusOK, gin.H { "todo": todo })
}

func deleteTodo(c *gin.Context) {
	Id := c.Param("ID")
	id, _ := strconv.Atoi(Id)

	db := gormConnect()

	var todo Todo

	db.Delete(&todo, id)
}
