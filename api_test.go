package main

import "testing"
import "github.com/stretchr/testify/assert"

func TestSaveAndFetchTask(t *testing.T) {
	db := NewSqlite("./gorm_test.db")
	api = API{db}
	defer api.Close()

	task := task{TaskID: 1, Title: "x", Content: "xx", Done: true}
	api.DeleteTask(task)
	api.CreateTask(task)

	maybeTask := api.FindByID(1)

	assert.Equal(t, task, maybeTask, "Can't save and fetch task")
}
