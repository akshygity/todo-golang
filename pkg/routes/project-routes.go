package routes

import (
	"github.com/akshay-attri/new/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterProjectRoutes = func(router *mux.Router) {
	router.HandleFunc("/Project/", controllers.CreateProject).Methods("POST")
	router.HandleFunc("/Project/", controllers.GetProject).Methods("GET")
	router.HandleFunc("/Project/{Title}", controllers.GetProjectByTitle).Methods("GET")
	router.HandleFunc("/Project/{ProjectId}", controllers.GetProjectById).Methods("GET")
	router.HandleFunc("/Projects/{ProjectId}", controllers.UpdateProject).Methods("PUT")
	router.HandleFunc("/Project/{Title}", controllers.UpdateProjectbyTitle).Methods("PUT")
	router.HandleFunc("/Project/{Title}/Archive", controllers.UpdateProjectbyArchive).Methods("PUT")
	router.HandleFunc("/Project/{ProjectId}", controllers.DeleteProject).Methods("DELETE")
}
