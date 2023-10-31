package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Smoked22/api-go-mysql/commons"
	"github.com/Smoked22/api-go-mysql/models"
	"github.com/gorilla/mux"
)

func GetAll(writer http.ResponseWriter, request *http.Request) {
	personas := []models.Persona{}
	db := commons.GetConnection()
	defer db.Close()

	db.Find(&personas)
	json, _ := json.Marshal(personas)
	commons.SendResponse(writer, http.StatusOK, json)
}

func Get(writer http.ResponseWriter, request *http.Request) {
	persona := models.Persona{}
	var result interface{}

	db := commons.GetConnection()
	defer db.Close()

	// Obtener el cuerpo del request (JSON)
	decoder := json.NewDecoder(request.Body)
	var filter map[string]interface{}
	err := decoder.Decode(&filter)
	if err != nil {
		commons.SendError(writer, http.StatusBadRequest)
		return
	}

	// Verificar si se proporcionó un ID o un nombre en el filtro
	if id, found := filter["id"]; found {
		db.Find(&persona, id)
		result = persona
	} else if nombre, found := filter["nombre"]; found {
		db.Where("nombre = ?", nombre).First(&persona)
		result = persona
	} else if apellido, found := filter["apellido"]; found {
		db.Where("apellido = ?", apellido).First(&persona)
		result = persona
	} else if direccion, found := filter["direccion"]; found {
		db.Where("direccion = ?", direccion).First(&persona)
		result = persona
	} else if telefono, found := filter["telefono"]; found {
		db.Where("telefono = ?", telefono).First(&persona)
		result = persona
	}

	if persona.ID > 0 {
		json, _ := json.Marshal(result)
		commons.SendResponse(writer, http.StatusOK, json)
	} else {
		commons.SendError(writer, http.StatusNotFound)
	}
}

func Save(writer http.ResponseWriter, request *http.Request) {
	persona := models.Persona{}

	db := commons.GetConnection()
	defer db.Close()

	error := json.NewDecoder(request.Body).Decode(&persona)

	if error != nil {
		log.Fatal(error)
		commons.SendError(writer, http.StatusBadRequest)
		return
	}

	error = db.Save(&persona).Error

	if error != nil {
		log.Fatal(error)
		commons.SendError(writer, http.StatusInternalServerError)
		return
	}

	json, _ := json.Marshal(persona)

	commons.SendResponse(writer, http.StatusCreated, json)
}

func Delete(writer http.ResponseWriter, request *http.Request) {
	persona := models.Persona{}

	db := commons.GetConnection()
	defer db.Close()

	id := mux.Vars(request)["id"]

	db.Find(&persona, id)

	if persona.ID > 0 {
		db.Delete(persona)
		commons.SendResponse(writer, http.StatusOK, []byte(`{}`))
	} else {
		commons.SendError(writer, http.StatusNotFound)
	}
}
