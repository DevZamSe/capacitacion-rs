package middleware

import (
	"capacitacion/src/entities"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func HolaMundo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	parametros := mux.Vars(r)

	nombre := parametros["nombre"]
	fmt.Println(nombre)

	json.NewEncoder(w).Encode(entities.ResponseHello{"Hola mundo", nombre})
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	fmt.Println(r.Header.Get("Authorization"))

	var datos entities.RequestName
	_ = json.NewDecoder(r.Body).Decode(&datos)

	fmt.Println(datos.Nombre)
	if datos.Nombre == "nilo" {
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(entities.ResponseGeneral{"login correcto"})
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(entities.ResponseGeneral{"fallo el login"})
	}
}
