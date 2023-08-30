package routes

import (
	"github.com/meech/auth_user_api/pkg/handlers"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/register", handlers.Register).Methods("POST")
	router.HandleFunc("/login", handlers.Login).Methods("POST")
	router.HandleFunc("/profile", handlers.UpdateProfile).Methods("PUT")
	router.HandleFunc("/delete", handlers.DeleteAccount).Methods("DELETE")
	router.HandleFunc("/users", handlers.GetAllUsers).Methods("GET")

	return router
}
