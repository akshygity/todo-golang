package models

import (
	"github.com/akshay-attri/new/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB


type Project struct {
	gorm.Model
	Title string `gorm:"unique" json:"title"`
	Archived bool `json:"archived"`
	Tasks []Task `gorm:"ForeignKey:ProjectID" json:"tasks"`
	}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Project{})
}

func (p *Project) CreateProject() *Project {
	db.NewRecord(p)
	db.Create(&p)
	return p

}

func Getallprojects() []Project {
	var Projects []Project
	db.Preload("Tasks").Find(&Projects)
	return Projects
	

}

func GetProjectbyId(Id int64) (*Project, *gorm.DB) {
	var getProject Project
	db := db.Where("ID=?", Id).Find(&getProject)
	return &getProject, db

}

func GetProjectByTitle(Title string )(*Project, *gorm.DB){
var getProject Project
db :=db.Where("Title=?", Title).Find(&getProject)
return &getProject,db
}


func GetTaskByProject(Title string ,Id int64)(*Project, *gorm.DB){
	var getProject Project
	db :=db.Where("Title=?", Title).Preload("Tasks").Where("ID=?", Id).Find(&getProject)
	//db :=db.Where("Title=?", Title).Find(&getProject)
	return &getProject,db
	}

	
	


func DeleteProject(ID int64) Project {
	var project Project
	db.Where("ID=?", ID).Delete(project)
	return project

}


