package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowIndexPage(ctx *gin.Context) {
	tasks := getAllTasks()
	ctx.HTML(http.StatusOK, "index.html", gin.H{"title": "To do list", "tasks": tasks})
}
