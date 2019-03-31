package main

func initRoutes() {
	router.GET("/", ShowIndexPage)
	router.POST("/newTask", CreateTask)
	router.POST("/updateTask", UpdateTask)
}
