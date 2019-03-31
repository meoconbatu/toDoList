package main

type FakeDatasource struct{}

func (fd FakeDatasource) GetAllTasks() []task {
	return []task{
		task{TaskID: 1, Title: "x", Content: "xx", Done: true},
		task{TaskID: 2, Title: "y", Content: "yy", Done: false},
	}
}

func (d FakeDatasource) CreateTask(t task) {
}

func (d FakeDatasource) UpdateTasks(tasks []task) {

}

func (d FakeDatasource) DeleteTask(t task) {

}
func (d FakeDatasource) FindByID(id int) task {
	return task{}
}
func (fd FakeDatasource) Close() {

}
