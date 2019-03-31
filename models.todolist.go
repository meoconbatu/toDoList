package main

type task struct {
	TaskID  uint16 `gorm:"primary_key:yes;column:taskID"`
	Title   string `form:"title" json:"title" xml:"title"  binding:"required"`
	Content string `form:"content" json:"content" xml:"content" binding:"required"`
	Done    bool   `form:"done" json:"done" xml:"done"`
}
