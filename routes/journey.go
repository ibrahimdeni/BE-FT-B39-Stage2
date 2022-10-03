package routes

import (
	"dumbflix/handlers"
	"dumbflix/pkg/middleware"
	"dumbflix/pkg/mysql"
	"dumbflix/repositories"

	"github.com/gorilla/mux"
)

func JourneyRoutes(r *mux.Router) {
	journeyRepository := repositories.RepositoryJourney(mysql.DB)
	h := handlers.HandlerJourney(journeyRepository)

	r.HandleFunc("/journeys", h.FindJourneys).Methods("GET")         //get alll
	r.HandleFunc("/journey/{id}", h.GetJourney).Methods("GET")       //select
	r.HandleFunc("/journey", middleware.Auth(middleware.UploadFile(h.CreateJourney))).Methods("POST")// add
}