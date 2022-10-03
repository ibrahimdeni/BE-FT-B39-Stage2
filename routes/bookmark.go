package routes

import (
	"dumbflix/handlers"
	"dumbflix/pkg/middleware"
	"dumbflix/pkg/mysql"
	"dumbflix/repositories"

	"github.com/gorilla/mux"
)

func BookmarkRoutes(r *mux.Router) {
	bookmarkRepository := repositories.RepositoryBookmark(mysql.DB)
	h := handlers.HandlerBookmark(bookmarkRepository)

	r.HandleFunc("/bookmarks", h.FindBookmarks).Methods("GET")         //get alll
	r.HandleFunc("/bookmark/{id}", h.GetBookmark).Methods("GET")       //select
	r.HandleFunc("/bookmark", middleware.Auth(h.CreateBookmark)).Methods("POST")// add
}