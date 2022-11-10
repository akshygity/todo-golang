package routes

import (
	"github.com/akshay-attri/new/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterTaskRoutes = func(router *mux.Router) {
	router.HandleFunc("/Task/", controllers.CreateTask).Methods("POST")
	router.HandleFunc("/Task/", controllers.GetTask).Methods("GET")
	router.HandleFunc("/Task/{TaskId}", controllers.GetTaskById).Methods("GET")
	router.HandleFunc("/Project/{Title}/tasks/{Id}", controllers.GetTaskByProject).Methods("GET")
	router.HandleFunc("/Project/{Title}/tasks/{Id}", controllers.UpdateProjectTask).Methods("PUT")
	router.HandleFunc("/Task/{TaskId}", controllers.UpdateTask).Methods("PUT")
	router.HandleFunc("/Project/{Title}/Tasks/", controllers.CreateTaskforProject).Methods("POST")
	router.HandleFunc("/Project/{Title}/Tasks/", controllers.GetProjectTasks).Methods("GET")
	router.HandleFunc("/Task/{TaskId}", controllers.DeleteTask).Methods("DELETE")
}
