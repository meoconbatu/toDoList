package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowIndexPage(ctx *gin.Context) {
	tasks := api.GetAllTasks()
	ctx.HTML(http.StatusOK, "index.html", gin.H{"title": "To do list", "tasks": tasks})
}

func CreateTask(c *gin.Context) {
	var t task
	if err := c.ShouldBind(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	api.CreateTask(t)

	c.Request.URL.Path = "/"
	c.Request.Method = "GET"
	router.HandleContext(c)
}
func UpdateTask(c *gin.Context) {
	doneTasks := c.PostFormArray("tasks")

	api.DataSource.UpdateTasks(doneTasks)

	c.Request.URL.Path = "/"
	c.Request.Method = "GET"
	router.HandleContext(c)
}
