package main

type DataSource interface {
	CreateTask(t task)
	UpdateTasks(tasks []task)
	DeleteTask(t task)

	FindByID(id int) task
	GetAllTasks() []task

	Close()
}
