package routes

import (
	"thejourney/handlers"
	"thejourney/pkg/middleware"
	"thejourney/pkg/mysql"
	"thejourney/repositories"

	"github.com/gorilla/mux"
)

func BookmarkRoutes(r *mux.Router) {
	bookmarkRepository := repositories.RepositoryBookmark(mysql.DB)
	h := handlers.HandlerBookmark(bookmarkRepository)

	r.HandleFunc("/bookmarks", h.FindBookmarks).Methods("GET")         //get alll
	r.HandleFunc("/bookmark/{id}", h.GetBookmark).Methods("GET")       //select
	r.HandleFunc("/bookmark", middleware.Auth(h.CreateBookmark)).Methods("POST")// add
	r.HandleFunc("/bookmark/{id}", middleware.Auth(h.DeleteBookmark)).Methods("DELETE") // delete
}