package main

type API struct {
	DataSource DataSource
}

func (api API) GetAllTasks() []task {
	return api.DataSource.GetAllTasks()
}
func (api API) CreateTask(t task) {
	api.DataSource.CreateTask(t)
}
func (api API) UpdateTasks(taskIDs []string) {
	api.DataSource.UpdateTasks(taskIDs)
}
func (api API) FindByID(id int) task {
	return api.DataSource.FindByID(id)
}
func (api API) DeleteTask(t task) {
	api.DataSource.DeleteTask(t)
}
func (api API) Close() {
	api.DataSource.Close()
}
