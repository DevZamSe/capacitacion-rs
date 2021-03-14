package routes

import (
	"capacitacion/src/middleware"

	"github.com/gorilla/mux"
)

func PublicWeb(r *mux.Router) {
	r.HandleFunc("/{nombre}", middleware.HolaMundo).Methods("GET")
	r.HandleFunc("/login", middleware.Login).Methods("POST")
}
