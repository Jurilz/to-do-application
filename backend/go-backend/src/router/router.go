package router

import (
	"github.com/gorilla/mux"
	"main/src/keycloak"
	"main/src/services"
	"net/http"
)

func CreateRouter(keycloak *keycloak.Keycloak) *mux.Router {

	router := mux.NewRouter()

	//router.HandleFunc("/tasks", services.GetTasks).Methods("GET")
	//router.HandleFunc("/tasks", services.CreateTask).Methods("POST")

	//router.HandleFunc("/tasks/{id}", services.GetTask).Methods("GET")
	//router.HandleFunc("/tasks/{id}", services.UpdateTask).Methods("PUT")
	//router.HandleFunc("/tasks/{id}", services.DeleteTask).Methods("DELETE")

	noAuthRouter := router.MatcherFunc(
		func(request *http.Request, routeMatcher *mux.RouteMatch) bool {
			return request.Header.Get("Authorization") == ""
		}).Subrouter()

	authRouter := router.MatcherFunc(
		func(request *http.Request, match *mux.RouteMatch) bool {
			return true
		}).Subrouter()

	controller := services.NewController(keycloak)

	noAuthRouter.HandleFunc(
		"/login", func(writer http.ResponseWriter, request *http.Request) {
			controller.Login(writer, request)
		}).Methods("POST")

	authRouter.HandleFunc("/tasks/{id}",
		services.GetTask).Methods("GET")

	authRouter.HandleFunc("/tasks",
		services.GetTasks).Methods("GET")

	authRouter.HandleFunc("/tasks",
		services.CreateTask).Methods("POST")

	authRouter.HandleFunc("/tasks/{id}",
		services.UpdateTask).Methods("PUT")

	authRouter.HandleFunc("/tasks/{id}",
		services.DeleteTask).Methods("DELETE")

	var middleware = services.NewMiddleware(keycloak)

	authRouter.Use(middleware.VerifyToken)

	return router
}
