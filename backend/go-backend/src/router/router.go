package router

import (
	"github.com/gorilla/mux"
	"main/src/services"
)

func CreateRouter() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/tasks", services.GetTasks).Methods("GET")
	router.HandleFunc("/tasks", services.CreateTask).Methods("POST")

	router.HandleFunc("/tasks/{id}", services.GetTask).Methods("GET")
	router.HandleFunc("/tasks/{id}", services.UpdateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id}", services.DeleteTask).Methods("DELETE")
	return router
}
