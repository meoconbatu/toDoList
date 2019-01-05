package main

import (
	"net/http"

	"github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
)

type Task struct {
	Title   string `form:"title" json:"title" xml:"title"  binding:"required"`
	Content string `form:"content" json:"content" xml:"content" binding:"required"`
}

var tasks []Task

func main() {
	router := gin.Default()
	router.HTMLRender = gintemplate.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "page.html", gin.H{"title": "To do list", "tasks": tasks})
	})
	router.POST("/newTask", func(c *gin.Context) {
		var t Task
		if err := c.ShouldBind(&t); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		tasks = append(tasks, t)
		c.Request.URL.Path = "/"
		c.Request.Method = "GET"
		router.HandleContext(c)
		//	c.JSON(http.StatusOK, gin.H{"status": "posted", "message": "title: " + t.Title + " content: " + t.Content})
	})
	router.Run(":8080")
}
