package main

import (
	"fmt"
	"net/http"

	"github.com/meech/auth_user_api/pkg/routes"

	"github.com/meech/auth_user_api/pkg/database"
)

func main() {
	database.Connect()

	router := routes.SetupRoutes()

	port := 8080
	fmt.Printf("Server started at :%d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), router)
}
