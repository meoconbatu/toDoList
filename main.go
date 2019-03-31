package main

import (
	"os"

	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
)

var (
	api    API
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

	db := NewSqlite("./gorm.db")
	api = API{db}
	defer api.Close()

	initRoutes()

	router.Run(":" + port)
}
