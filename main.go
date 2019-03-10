package main

import (
	"fmt"
	"os"

	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Task struct {
	TaskID  uint16 `gorm:"primary_key:yes;column:taskID"`
	Title   string `form:"title" json:"title" xml:"title"  binding:"required"`
	Content string `form:"content" json:"content" xml:"content" binding:"required"`
	Done    bool   `form:"done" json:"done" xml:"done"`
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

	initRoutes()

	router.Run(":8080")
}
