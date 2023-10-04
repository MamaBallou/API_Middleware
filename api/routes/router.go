package routes

import (
	"middleware/api/controllers"
	"middleware/api/middleware_api"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/mux"
)

// Crée un nouveau routeur en utilisant le package mux.
func NewRouter() *mux.Router {
    r := mux.NewRouter().StrictSlash(true)

    // Ajoute des middlewares globaux.
    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)
	r.Use(middleware_api.JwtMiddleware)
    // Charge la gestion des en-têtes CORS.
    LoadCors(r)

    // Configure les routes pour diverses actions avec leurs contrôleurs associés.

	r.HandleFunc("/middleware", controllers.MiddlewareGet).Methods("GET")
	r.HandleFunc("/middleware", controllers.MiddlewarePost).Methods("POST")
	r.HandleFunc("/middleware", controllers.MiddlewarePatch).Methods("PATCH")
	r.HandleFunc("/middleware", controllers.MiddlewareDelete).Methods("DELETE")

	return r
}