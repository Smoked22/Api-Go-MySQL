package routes

import (
	"github.com/Smoked22/api-go-mysql/controllers"
	"github.com/gorilla/mux"
)

func SetPersonaRoutes(router *mux.Router) {
	subRoute := router.PathPrefix("/persona/api").Subrouter()

	subRoute.HandleFunc("/all", controllers.GetAll).Methods("GET")
	subRoute.HandleFunc("/save", controllers.Save).Methods("POST")
	subRoute.HandleFunc("/delete/{id}", controllers.Delete).Methods("POST")
	subRoute.HandleFunc("/find", controllers.Get).Methods("GET")

}
