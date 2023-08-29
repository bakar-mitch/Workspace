package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/meech/task_manager/pkg/controllers"
	"github.com/meech/task_manager/pkg/database"
)

func main() {
	database.Connect()
	r := mux.NewRouter()

	r.HandleFunc("/tasks/create", controllers.CreateTask).Methods("POST")
	r.HandleFunc("/tasks/{id}", controllers.GetTask).Methods("GET")
	r.HandleFunc("/tasks", controllers.GetAllTasks).Methods("GET")
	r.HandleFunc("/tasks/{id}", controllers.UpdateTask).Methods("PUT")
	r.HandleFunc("/tasks/{id}", controllers.DeleteTask).Methods("DELETE")
	// Define routes for UpdateTask and DeleteTask

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
