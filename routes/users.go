package routes

import (
	"dumbflix/handlers"
	"dumbflix/pkg/mysql"
	"dumbflix/repositories"

	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router) {
	userRepository := repositories.RepositoryUser(mysql.DB)
	h := handlers.HandlerUser(userRepository)

	r.HandleFunc("/users", h.FindUsers).Methods("GET")         //get alll
	r.HandleFunc("/user/{id}", h.GetUser).Methods("GET")       //select
	// r.HandleFunc("/user", h.CreateUser).Methods("POST")        // add
	r.HandleFunc("/user/{id}", h.UpdateUser).Methods("PATCH")  // edite
	r.HandleFunc("/user/{id}", h.DeleteUser).Methods("DELETE") // delete
}
