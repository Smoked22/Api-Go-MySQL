package main

import (
	"log"
	"net/http"

	"github.com/Smoked22/api-go-mysql/commons"
	"github.com/Smoked22/api-go-mysql/routes"
	"github.com/gorilla/mux"
)

func main() {
	commons.Migrate()

	router := mux.NewRouter()
	routes.SetPersonaRoutes(router)

	server := http.Server{
		Addr:    ":9000",
		Handler: router,
	}

	log.Println("Servidor ejecutandose sobre el puerto 9000")

	log.Println(server.ListenAndServe())
}
