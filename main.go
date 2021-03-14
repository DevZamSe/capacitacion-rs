package main

import (
	"capacitacion/routes"
	"capacitacion/src/config"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {

	config.GetMySQLDB()

	Routes()
}

func Routes() *mux.Router {

	r := mux.NewRouter()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*", "http://localhost:8000"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"*", "Authorization", "Content-Type", "Access-Control-Allow-Origin"},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
			http.MethodHead,
		},
		Debug: true,
	})

	routes.PublicWeb(r)

	fmt.Println("abierto correctamente")
	http.ListenAndServe(":8000", c.Handler(r))

	return r
}
