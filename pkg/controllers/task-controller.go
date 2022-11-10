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

var NewTask models.Task

func GetTask(w http.ResponseWriter, r *http.Request) {
	newTasks := models.Getalltasks()
	res, _ := json.Marshal(newTasks)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetTaskById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskId := vars["TaskId"]
	ID, err := strconv.ParseInt(taskId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	taskDetails, _ := models.GetTaskbyId(ID)
	res, _ := json.Marshal(taskDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


func GetProjectTasks(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	projectTitle := vars["Title"]
	
	projectDetails, _ := models.GetProjectByTitle(projectTitle)
	projectId :=projectDetails.ID
	taskDetails:= models.GetTaskProject(projectId)
	res, _ := json.Marshal(taskDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}



func CreateTask(w http.ResponseWriter, r *http.Request) {
	CreateTask := &models.Task{}
	utils.Parsebody(r, CreateTask)
	b := CreateTask.CreateTask()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "pkglication/json")
	w.Write(res)

}
func CreateTaskforProject(w http.ResponseWriter, r *http.Request) {
	CreateTask := &models.Task{}
	vars := mux.Vars(r)
	projectTitle := vars["Title"]
	
	projectDetails, _ := models.GetProjectByTitle(projectTitle)
	CreateTask.ProjectID=projectDetails.ID
	utils.Parsebody(r, CreateTask)
	b := CreateTask.CreateTask()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "pkglication/json")
	w.Write(res)

}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskId := vars["TaskId"]
	ID, err := strconv.ParseInt(taskId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	task := models.DeleteTask(ID)
	res, _ := json.Marshal(task)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	var updateTask = &models.Task{}
	utils.Parsebody(r, updateTask)
	vars := mux.Vars(r)
	taskId := vars["TaskId"]
	ID, err := strconv.ParseInt(taskId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	taskDetails, db := models.GetTaskbyId(ID)
	if updateTask.Title != "" {
		taskDetails.Title = updateTask.Title
	}
	if updateTask.Priority != "" {
		taskDetails.Priority = updateTask.Priority
	}
	// if updateTask.Task != "" {
	// 	taskDetails.Task = updateTask.Task
	// }
	db.Save(&taskDetails)
	res, _ := json.Marshal(taskDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

// func UpdateTaskByProject(w http.ResponseWriter, r *http.Request) {
// 	var updateTask = &models.Task{}
// 	utils.Parsebody(r, updateTask)
// 	vars := mux.Vars(r)
// 	projectTitle :=vars["Title"]
// 	TaskId :=vars["Id"]
// 	Id, err := strconv.ParseInt(TaskId, 0, 0)
// 	if err != nil {
// 		fmt.Println("Error while parsing")
// 	}
// 	taskDetails, db := models.GetTaskProject(projectTitle,Id)
// 	if updateTask.Title != "" {
// 		taskDetails.Title = updateTask.Title
// 	}
// 	if updateTask.Priority != "" {
// 		taskDetails.Priority = updateTask.Priority
// 	}
// 	if updateTask.Deadline != nil  {
// 		taskDetails.Deadline = updateTask.Deadline
// 	}
// 	db.Save(&taskDetails)
// 	res, _ := json.Marshal(taskDetails)
// 	w.Header().Set("Content-Type", "pkglication/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)

// }



func GetTaskByProject(w http.ResponseWriter ,r *http.Request){
	vars :=mux.Vars(r)
	//projectTitle :=vars["Title"]
	TaskId :=vars["Id"]
	Id, err := strconv.ParseInt(TaskId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	taskDetails, _ := models.GetTaskbyId(Id)
	res, _ := json.Marshal(taskDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
 
 
 
 }
 
 func UpdateProjectTask(w http.ResponseWriter, r *http.Request) {
	var updateTask = &models.Task{}
	utils.Parsebody(r, updateTask)
	vars := mux.Vars(r)
	taskId := vars["TaskId"]
	Id, err := strconv.ParseInt(taskId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	
	taskDetails, db := models.GetTaskbyId(Id)
	if taskDetails.Title != ""{
		taskDetails.Title =updateTask.Title
	}
	if taskDetails.Priority != ""{
		taskDetails.Priority =updateTask.Priority
	}

		taskDetails.Deadline =updateTask.Deadline
	
	
		taskDetails.Done =updateTask.Done
	
	
		taskDetails.ProjectID =updateTask.ProjectID
	

	db.Save(&taskDetails)
	res, _ := json.Marshal(taskDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}