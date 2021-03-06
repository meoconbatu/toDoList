package main

import (
	"fmt"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var taskList []task

type Sqlite struct {
	db *gorm.DB
}

func NewSqlite(args string) Sqlite {
	db, err := gorm.Open("sqlite3", args)
	if err != nil {
		fmt.Println(err)
	}
	db.AutoMigrate(&task{})
	return Sqlite{db}
}

func (s Sqlite) GetAllTasks() []task {
	s.db.Find(&taskList)
	return taskList
}

func (s Sqlite) CreateTask(t task) {
	s.db.Create(&t)
}

func (s Sqlite) UpdateTasks(tasks []task) {
	for _, t := range tasks {
		s.db.Save(&t)
	}
}
func (s Sqlite) FindByID(id int) task {
	var t task
	s.db.Where("taskid = ?", id).First(&t)
	return t
}
func (s Sqlite) DeleteTask(t task) {
	s.db.Delete(&t)
}
func (s Sqlite) Close() {
	s.db.Close()
}
