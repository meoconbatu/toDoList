package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initRoutes() {
	router.GET("/", ViewTask)
	router.POST("/newTask", CreateTask)
	router.POST("/updateTask", UpdateTask)
}
func ViewTask(ctx *gin.Context) {
	var tasks []Task
	db.Find(&tasks)
	ctx.HTML(http.StatusOK, "page.html", gin.H{"title": "To do list", "tasks": tasks})
}
func CreateTask(c *gin.Context) {
	var t Task
	if err := c.ShouldBind(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&t)
	c.Request.URL.Path = "/"
	c.Request.Method = "GET"
	router.HandleContext(c)
}
func UpdateTask(c *gin.Context) {
	var tasks []Task
	doneTasks := c.PostFormArray("tasks")

	db.Find(&tasks)
	db.Table("tasks").Update("done", "false")

	for _, t := range doneTasks {
		var task Task
		db.Where("taskid = ?", t).First(&task)
		db.Model(&task).Update("done", true)
	}

	c.Request.URL.Path = "/"
	c.Request.Method = "GET"
	router.HandleContext(c)
}
