package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initRoutes() {
	router.GET("/", ShowIndexPage)
	router.POST("/newTask", CreateTask)
	router.POST("/updateTask", UpdateTask)
}
func CreateTask(c *gin.Context) {
	var t task
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
	var tasks []task
	doneTasks := c.PostFormArray("tasks")

	db.Find(&tasks)
	db.Table("tasks").Update("done", "false")

	for _, t := range doneTasks {
		var task task
		db.Where("taskid = ?", t).First(&task)
		db.Model(&task).Update("done", true)
	}

	c.Request.URL.Path = "/"
	c.Request.Method = "GET"
	router.HandleContext(c)
}
