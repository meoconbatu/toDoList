package main

import (
	"fmt"
	"os"

	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	db     *gorm.DB
	err    error
	router *gin.Engine
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		// log.Fatal("$PORT must be set")
	}
	router = gin.Default()
	router.HTMLRender = gintemplate.Default()

	// dialect := os.Getenv("TODOLIST_DIALECT") //"sqlite3"
	// args := os.Getenv("TODOLIST_PARAM")      //"./gorm.db"
	// db, err = gorm.Open(dialect, args)
	db, err = gorm.Open("sqlite3", "./gorm.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	db.AutoMigrate(&task{})
	initRoutes()

	// router.Run(":8080")
	router.Run(":" + port)
}
