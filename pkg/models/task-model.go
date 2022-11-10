	package models

import (
	"github.com/akshay-attri/new/pkg/config"
	"github.com/jinzhu/gorm"
	"time"
)



type Task struct {
	gorm.Model
	Title string `json:"title"`
	Priority string `gorm:"type:ENUM('0', '1', '2', '3');default:'0'" json:"priority"`
	Deadline *time.Time `gorm:"default:null" json:"deadline"`
	Done bool `json:"done"`
	ProjectID uint `json:"project_id"`
	}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Task{})
}

func (t *Task) CreateTask() *Task {
	db.NewRecord(t)
	db.Create(&t)
	return t

}

func Getalltasks() []Task {
	var Tasks []Task
	db.Find(&Tasks)
	return Tasks

}

func GetTaskbyId(Id int64) (*Task, *gorm.DB) {
	var getTask Task
	db := db.Where("ID=?", Id).Find(&getTask)
	return &getTask, db

}

func DeleteTask(ID int64) Task {
	var task Task
	db.Where("ID=?", ID).Delete(task)
	return task

}
func GetTaskProject(ProjectId uint)[]Task{
	var getTask []Task
   db.Where("project_Id=?", ProjectId).Find(&getTask)
	return getTask
	}