package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/akshay-attri/new/pkg/models"
	"github.com/akshay-attri/new/pkg/utils"

	"github.com/gorilla/mux"
)

var NewProject models.Project

func GetProject(w http.ResponseWriter, r *http.Request) {
	newProjects := models.Getallprojects()
	res, _ := json.Marshal(newProjects)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)

	w.Write(res)
}

func GetProjectById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	projectId := vars["ProjectId"]
	ID, err := strconv.ParseInt(projectId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	projectDetails, _ := models.GetProjectbyId(ID)
	res, _ := json.Marshal(projectDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


func GetProjectByTitle(w http.ResponseWriter ,r *http.Request){
   vars :=mux.Vars(r)
   projectTitle :=vars["Title"]
   projectDetails, _ := models.GetProjectByTitle(projectTitle)
   res, _ := json.Marshal(projectDetails)
   w.Header().Set("Content-Type", "pkglication/json")
   w.WriteHeader(http.StatusOK)
   w.Write(res)



}



func CreateProject(w http.ResponseWriter, r *http.Request) {
	CreateProject := &models.Project{}
	utils.Parsebody(r, CreateProject)
	b := CreateProject.CreateProject()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "pkglication/json")
	w.Write(res)

}

func DeleteProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	projectId := vars["ProjectId"]
	ID, err := strconv.ParseInt(projectId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	project := models.DeleteProject(ID)
	res, _ := json.Marshal(project)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func UpdateProject(w http.ResponseWriter, r *http.Request) {
	var updateProject = &models.Project{}
	utils.Parsebody(r, updateProject)
	vars := mux.Vars(r)
	projectId := vars["projectId"]
	ID, err := strconv.ParseInt(projectId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	projectDetails, db := models.GetProjectbyId(ID)
	if updateProject.Title != "" {
		projectDetails.Title = updateProject.Title
	}
	if updateProject.Archived != false {
		projectDetails.Archived = updateProject.Archived
	}
	
	db.Save(&projectDetails)
	res, _ := json.Marshal(projectDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


func UpdateProjectbyTitle(w http.ResponseWriter, r *http.Request) {
	var updateProject = &models.Project{}
	utils.Parsebody(r, updateProject)
	vars := mux.Vars(r)
	projectTitle := vars["Title"]
	
	projectDetails, db := models.GetProjectByTitle(projectTitle)
	if updateProject.Title != "" {
		projectDetails.Title = updateProject.Title
	}
	if updateProject.Archived != false {
		projectDetails.Archived = updateProject.Archived
	}
	
	db.Save(&projectDetails)
	res, _ := json.Marshal(projectDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
func UpdateProjectbyArchive(w http.ResponseWriter, r *http.Request) {
	var updateProject = &models.Project{}
	utils.Parsebody(r, updateProject)
	vars := mux.Vars(r)
	projectTitle := vars["Title"]
	
	projectDetails, db := models.GetProjectByTitle(projectTitle)
	

		projectDetails.Archived = updateProject.Archived
	
	
	db.Save(&projectDetails)
	res, _ := json.Marshal(projectDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}


