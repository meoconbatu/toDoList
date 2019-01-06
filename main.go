package main

import (
	"fmt"
	"net/http"

	"github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Task struct {
	Title   string `form:"title" json:"title" xml:"title"  binding:"required"`
	Content string `form:"content" json:"content" xml:"content" binding:"required"`
}

var (
	db  *gorm.DB
	err error
)

func main() {
	router := gin.Default()
	router.HTMLRender = gintemplate.Default()

	db, err = gorm.Open("sqlite3", "./gorm.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	db.AutoMigrate(&Task{})

	router.GET("/", func(ctx *gin.Context) {
		var tasks []Task
		db.Find(&tasks)
		ctx.HTML(http.StatusOK, "page.html", gin.H{"title": "To do list", "tasks": tasks})
	})

	router.POST("/newTask", func(c *gin.Context) {
		var t Task
		if err := c.ShouldBind(&t); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Create(&t)
		c.Request.URL.Path = "/"
		c.Request.Method = "GET"
		router.HandleContext(c)
	})

	router.Run(":8080")
}
