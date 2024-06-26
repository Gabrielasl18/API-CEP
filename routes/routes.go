package routes

import (
	"modules-app/controllers"

	"github.com/gorilla/mux"
)

func SetupRoutes(controller *controllers.Controller) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", controller.GetCEPS).Methods("GET")
	router.HandleFunc("/{id}", controller.GetCEPByID).Methods("GET")
	router.HandleFunc("/create/cep", controller.CreateCEP).Methods("POST")

	return router
}
