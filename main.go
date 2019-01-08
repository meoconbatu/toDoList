package main

import (
	"fmt"
	"net/http"
	"os"

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
	db     *gorm.DB
	err    error
	router *gin.Engine
)

func main() {
	router = gin.Default()
	router.HTMLRender = gintemplate.Default()

	dialect := os.Getenv("TODOLIST_DIALECT")
	param := os.Getenv("TODOLIST_PARAM")
	// db, err = gorm.Open("sqlite3", "./gorm.db")
	db, err = gorm.Open(dialect, param)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	db.AutoMigrate(&Task{})

	router.GET("/", ViewTask)
	router.POST("/newTask", CreateTask)

	router.Run(":8080")
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
