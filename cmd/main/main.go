package main

import (
	"log"
	"net/http"

	"github.com/akshay-attri/new/pkg/routes"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()

	routes.RegisterProjectRoutes(r)
	routes.RegisterTaskRoutes(r)
	
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}






