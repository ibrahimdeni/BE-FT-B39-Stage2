package routes

import (
	"thejourney/handlers"
	"thejourney/pkg/middleware"
	"thejourney/pkg/mysql"
	"thejourney/repositories"

	"github.com/gorilla/mux"
)

func JourneyRoutes(r *mux.Router) {
	journeyRepository := repositories.RepositoryJourney(mysql.DB)
	h := handlers.HandlerJourney(journeyRepository)

	r.HandleFunc("/journeys", h.FindJourneys).Methods("GET")         //get alll
	r.HandleFunc("/journey/{id}", h.GetJourney).Methods("GET")       //select
	r.HandleFunc("/journey", middleware.Auth(middleware.UploadFile(h.CreateJourney))).Methods("POST")// add
	r.HandleFunc("/journey/{id}", middleware.Auth(middleware.UploadFile(h.UpdateJourney))).Methods("PATCH")  // edite
	r.HandleFunc("/journey/{id}", middleware.Auth(h.DeleteJourney)).Methods("DELETE") // delete
}