package main

type DataSource interface {
	CreateTask(t task)
	UpdateTasks(taskIDs []string)
	DeleteTask(t task)

	FindByID(id int) task
	GetAllTasks() []task

	Close()
}
